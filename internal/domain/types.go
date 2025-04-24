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

type Order struct {
	ID            string      `json:"id"`
	PaymentMethod string      `json:"payment_method"`
	TaxPrice      float64     `json:"tax_price"`
	ShippingPrice float64     `json:"shipping_price"`
	TotalPrice    float64     `json:"total_price"`
	OrderItems    []OrderItem `json:"order_items"`
	CreatedAt     uint64      `json:"created_at"`
	UpdatedAt     uint64      `json:"updated_at"`
}

type OrderItem struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Image     string  `json:"image"`
	Price     float64 `json:"price"`
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
	Name            string  `json:"name"`
	Image           string  `json:"image"`
	Category        string  `json:"category"`
	Description     string  `json:"description"`
	Rating          int     `json:"rating"`
	NumberOfReviews int     `json:"number_of_reviews"`
	Price           float64 `json:"price"`
	CountInStock    int     `json:"count_in_stock"`
}

type CreateOrderRequest struct {
	PaymentMethod string      `json:"payment_method" binding:"required"`
	TaxPrice      float64     `json:"tax_price" binding:"required"`
	ShippingPrice float64     `json:"shipping_price" binding:"required"`
	TotalPrice    float64     `json:"total_price" binding:"required"`
	OrderItems    []OrderItem `json:"order_items" binding:"required"`
}
