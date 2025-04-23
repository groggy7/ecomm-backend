package domain

type Product struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Image           string  `json:"image"`
	Category        string  `json:"category"`
	Description     string  `json:"description"`
	Rating          int     `json:"rating"`
	NumberOfReviews int     `json:"number_of_reviews"`
	Price           float64 `json:"price"`
	CountInStock    int     `json:"count_in_stock"`
	CreatedAt       uint64  `json:"created_at"`
	UpdatedAt       uint64  `json:"updated_at"`
}

type CreateProductRequest struct {
	Name            string  `json:"name" binding:"required"`
	Image           string  `json:"image" binding:"required"`
	Category        string  `json:"category" binding:"required"`
	Description     string  `json:"description" binding:"required"`
	Rating          int     `json:"rating" binding:"required"`
	NumberOfReviews int     `json:"number_of_reviews" binding:"required"`
	Price           float64 `json:"price" binding:"required"`
	CountInStock    int     `json:"count_in_stock" binding:"required"`
}

type UpdateProductRequest struct {
	ID              string
	Name            string  `json:"name" binding:"required"`
	Image           string  `json:"image" binding:"required"`
	Category        string  `json:"category" binding:"required"`
	Description     string  `json:"description" binding:"required"`
	Rating          int     `json:"rating" binding:"required"`
	NumberOfReviews int     `json:"number_of_reviews" binding:"required"`
	Price           float64 `json:"price" binding:"required"`
	CountInStock    int     `json:"count_in_stock" binding:"required"`
}

type ProductRepository interface {
	CreateProduct(product *Product) (*Product, error)
	GetProductByID(id string) (*Product, error)
	GetAllProducts() ([]Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id string) error
}
