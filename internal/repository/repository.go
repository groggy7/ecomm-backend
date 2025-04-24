package repository

import (
	"context"
	"ecomm/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) domain.Repository {
	return &repository{pool: pool}
}

func (r *repository) CreateProduct(product *domain.Product) (*domain.Product, error) {
	query := `
        INSERT INTO 
        products(name, image, category, description, rating, num_reviews, price, count_in_stock)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
    `

	err := r.pool.QueryRow(context.Background(), query,
		&product.Name,
		&product.Image,
		&product.Category,
		&product.Description,
		&product.Rating,
		&product.NumberOfReviews,
		&product.Price,
		&product.CountInStock).Scan(&product.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repository) GetProductByID(id string) (*domain.Product, error) {
	query := `
		SELECT id, name, image, category, description, rating, num_reviews, price, count_in_stock, created_at, updated_at
		FROM products WHERE id = $1
	`

	product := new(domain.Product)
	if err := r.pool.QueryRow(context.Background(), query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Image,
		&product.Category,
		&product.Description,
		&product.Rating,
		&product.NumberOfReviews,
		&product.Price,
		&product.CountInStock,
		&product.CreatedAt,
		&product.UpdatedAt); err != nil {
		return nil, err
	}

	return product, nil
}

func (r *repository) ListProducts() ([]domain.Product, error) {
	query := `
		SELECT id, name, image, category, description, rating, num_reviews, price, count_in_stock, created_at, updated_at
		FROM products
	`

	products := make([]domain.Product, 0)
	rows, err := r.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := new(domain.Product)
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Image,
			&product.Category,
			&product.Description,
			&product.Rating,
			&product.NumberOfReviews,
			&product.Price,
			&product.CountInStock,
			&product.CreatedAt,
			&product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}

func (r *repository) UpdateProduct(product *domain.Product) error {
	query := `
		UPDATE products
		SET name = $1, image = $2, category = $3, description = $4,
		rating = $5, num_reviews = $6, price = $7, count_in_stock = $8
		WHERE id = $9
	`

	if _, err := r.pool.Exec(context.Background(), query,
		&product.Name,
		&product.Image,
		&product.Category,
		&product.Description,
		&product.Rating,
		&product.NumberOfReviews,
		&product.Price,
		&product.CountInStock,
		&product.ID); err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteProduct(id string) error {
	query := `DELETE FROM products where id = $1`
	result, err := r.pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrProductNotFound
	}

	return nil
}

func (r *repository) CreateOrder(order *domain.Order) (*domain.Order, error) {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(context.Background())

	query := `
		INSERT INTO orders(payment_method, tax_price, shipping_price, total_price)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err = tx.QueryRow(context.Background(), query,
		&order.PaymentMethod,
		&order.TaxPrice,
		&order.ShippingPrice,
		&order.TotalPrice).Scan(&order.ID)
	if err != nil {
		return nil, err
	}

	query = `
		INSERT INTO order_items(order_id, product_id, name, quantity, image, price)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	for _, orderItem := range order.OrderItems {
		err := tx.QueryRow(context.Background(), query,
			&order.ID,
			&orderItem.ProductID,
			&orderItem.Name,
			&orderItem.Quantity,
			&orderItem.Image,
			&orderItem.Price).Scan(&orderItem.ID)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return order, nil
}

func (r *repository) GetOrderByID(id string) (*domain.Order, error) {
	query := `
		SELECT id, payment_method, tax_price, shipping_price, total_price, created_at, updated_at
		FROM orders WHERE id = $1
	`

	order := new(domain.Order)
	if err := r.pool.QueryRow(context.Background(), query, id).Scan(
		&order.ID,
		&order.PaymentMethod,
		&order.TaxPrice,
		&order.ShippingPrice,
		&order.TotalPrice,
		&order.CreatedAt,
		&order.UpdatedAt); err != nil {
		return nil, err
	}

	return order, nil
}

func (r *repository) ListOrders() ([]domain.Order, error) {
	query := `
		SELECT id, payment_method, tax_price, shipping_price, total_price, created_at, updated_at
		FROM orders
	`

	var orders []domain.Order
	if err := pgxscan.Select(context.Background(), r.pool, &orders, query); err != nil {
		return nil, err
	}

	for i := range orders {
		query := `
			SELECT id, order_id, product_id, name, quantity, image, price
			FROM order_items WHERE order_id = $1
		`

		var orderItems []domain.OrderItem
		if err := pgxscan.Select(context.Background(), r.pool, &orderItems, query, orders[i].ID); err != nil {
			return nil, err
		}

		orders[i].OrderItems = orderItems
	}

	return orders, nil
}

func (r *repository) DeleteOrder(id string) error {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	query := `DELETE FROM order_items where order_id = $1`
	if _, err := tx.Exec(context.Background(), query, id); err != nil {
		return err
	}

	query = `DELETE FROM orders where id = $1`
	result, err := tx.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrOrderNotFound
	}

	if err := tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetOrderItems(orderID string) ([]domain.OrderItem, error) {
	query := `
		SELECT id, order_id, product_id, name, quantity, image, price
		FROM order_items WHERE order_id = $1
	`

	var orderItems []domain.OrderItem
	if err := pgxscan.Select(context.Background(), r.pool, &orderItems, query, orderID); err != nil {
		return nil, err
	}

	return orderItems, nil
}
