package domain

type Repository interface {
	CreateProduct(product *Product) (*Product, error)
	GetProductByID(id string) (*Product, error)
	ListProducts() ([]Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id string) error

	CreateOrder(order *Order) (*Order, error)
	GetOrderByID(id string) (*Order, error)
	ListOrders() ([]Order, error)
	DeleteOrder(id string) error
}
