package domain

type Repository interface {
	CreateProduct(product *Product) (*Product, error)
	GetProductByID(id string) (*Product, error)
	ListProducts() ([]*Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id string) error

	CreateOrder(order *Order) (*Order, error)
	GetOrder(userID string) (*Order, error)
	GetOrderItems(orderID string) ([]OrderItem, error)
	ListOrders() ([]*Order, error)
	DeleteOrder(id string) error

	CreateUser(user *User) (*User, error)
	GetUser(email string) (*User, error)
	ListUsers() ([]*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error

	CreateSession(session *Session) error
	GetSession(id string) (*Session, error)
	RevokeSession(id string) error
	DeleteSession(id string) error
}
