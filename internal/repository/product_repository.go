package repository

import (
	"context"
	"ecomm/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepo struct {
	pool *pgxpool.Pool
}

func NewProductRepo(pool *pgxpool.Pool) domain.ProductRepository {
	return &productRepo{pool: pool}
}

func (p *productRepo) CreateProduct(product *domain.Product) (*domain.Product, error) {
	query := `
        INSERT INTO 
        products(name, image, category, description, rating, num_reviews, price, count_in_stock)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
    `

	err := p.pool.QueryRow(context.Background(), query,
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

func (p *productRepo) GetProductByID(id string) (*domain.Product, error) {
	query := `
		SELECT id, name, image, category, description, rating, num_reviews, price, count_in_stock, created_at, updated_at
		FROM products WHERE id = $1
	`

	product := new(domain.Product)
	if err := p.pool.QueryRow(context.Background(), query, id).Scan(
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

func (p *productRepo) GetAllProducts() ([]domain.Product, error) {
	query := `
		SELECT id, name, image, category, description, rating, num_reviews, price, count_in_stock, created_at, updated_at
		FROM products
	`

	products := make([]domain.Product, 0)
	rows, err := p.pool.Query(context.Background(), query)
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

func (p *productRepo) UpdateProduct(product *domain.Product) error {
	query := `
		UPDATE products
		SET name = $1, image = $2, category = $3, description = $4,
		rating = $5, num_reviews = $6, price = $7, count_in_stock = $8
		WHERE id = $9
	`

	if _, err := p.pool.Exec(context.Background(), query,
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

func (p *productRepo) DeleteProduct(id string) error {
	query := `DELETE FROM products where id = $1`
	if _, err := p.pool.Exec(context.Background(), query, id); err != nil {
		return err
	}

	return nil
}
