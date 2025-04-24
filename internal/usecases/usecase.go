package usecases

import "ecomm/internal/domain"

type UseCase struct {
	repo domain.Repository
}

func NewUseCase(repo domain.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) CreateProduct(req *domain.CreateProductRequest) (*domain.Product, error) {
	product := &domain.Product{
		Name:            req.Name,
		Image:           req.Image,
		Category:        req.Category,
		Description:     req.Description,
		Rating:          req.Rating,
		NumberOfReviews: req.NumberOfReviews,
		Price:           req.Price,
		CountInStock:    req.CountInStock,
	}

	return u.repo.CreateProduct(product)
}

func (u *UseCase) GetProductByID(id string) (*domain.Product, error) {
	return u.repo.GetProductByID(id)
}

func (u *UseCase) GetAllProducts() ([]domain.Product, error) {
	return u.repo.ListProducts()
}

func (u *UseCase) UpdateProduct(req *domain.UpdateProductRequest) error {
	product, err := u.repo.GetProductByID(req.ID)
	if err != nil {
		return err
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Image != "" {
		product.Image = req.Image
	}
	if req.Category != "" {
		product.Category = req.Category
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Rating != 0 {
		product.Rating = req.Rating
	}
	if req.NumberOfReviews != 0 {
		product.NumberOfReviews = req.NumberOfReviews
	}
	if req.Price != 0 {
		product.Price = req.Price
	}
	if req.CountInStock != 0 {
		product.CountInStock = req.CountInStock
	}

	return u.repo.UpdateProduct(product)
}

func (u *UseCase) DeleteProduct(id string) error {
	return u.repo.DeleteProduct(id)
}

func (u *UseCase) CreateOrder(req *domain.CreateOrderRequest) (*domain.Order, error) {
	for _, item := range req.OrderItems {
		_, err := u.repo.GetProductByID(item.ProductID)
		if err != nil {
			return nil, err
		}
	}

	order := &domain.Order{
		PaymentMethod: req.PaymentMethod,
		TaxPrice:      req.TaxPrice,
		ShippingPrice: req.ShippingPrice,
		TotalPrice:    req.TotalPrice,
		OrderItems:    req.OrderItems,
	}

	return u.repo.CreateOrder(order)
}

func (u *UseCase) GetOrderByID(id string) (*domain.Order, error) {
	order, err := u.repo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	orderItems, err := u.repo.GetOrderItems(id)
	if err != nil {
		return nil, err
	}

	order.OrderItems = orderItems
	return order, nil
}

func (u *UseCase) GetAllOrders() ([]domain.Order, error) {
	orders, err := u.repo.ListOrders()
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return []domain.Order{}, nil
	}

	return orders, nil
}

func (u *UseCase) DeleteOrder(id string) error {
	return u.repo.DeleteOrder(id)
}
