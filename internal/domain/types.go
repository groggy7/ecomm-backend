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
	Name            string  `json:"name"`
	Image           string  `json:"image"`
	Category        string  `json:"category"`
	Description     string  `json:"description"`
	Rating          int     `json:"rating"`
	NumberOfReviews int     `json:"number_of_reviews"`
	Price           float64 `json:"price"`
	CountInStock    int     `json:"count_in_stock"`
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

type CreateOrderRequest struct {
	PaymentMethod string      `json:"payment_method" binding:"required"`
	TaxPrice      float64     `json:"tax_price" binding:"required"`
	ShippingPrice float64     `json:"shipping_price" binding:"required"`
	TotalPrice    float64     `json:"total_price" binding:"required"`
	OrderItems    []OrderItem `json:"order_items" binding:"required"`
}

type OrderItem struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id" binding:"required"`
	Name      string  `json:"name"  binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	Image     string  `json:"image" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"updated_at"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsAdmin  bool   `json:"is_admin"`
}

type CreateUserResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

type ListUserResponse struct {
	Users []UserInfo `json:"users"`
}

type UserInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type Session struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	RefreshToken string `json:"refresh_token"`
	IsRevoked    bool   `json:"is_revoked"`
	CreatedAt    uint64 `json:"created_at"`
	ExpiresAt    uint64 `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	SessionID             string `json:"session_id"`
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AccessTokenExpiresAt  uint64 `json:"access_token_expires_at"`
	RefreshTokenExpiresAt uint64 `json:"refresh_token_expires_at"`
}

type LogoutRequest struct {
	Email     string `json:"email" binding:"required"`
	SessionID string `json:"session_id" binding:"required"`
}

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
	SessionID    string `json:"session_id" binding:"required"`
}

type RefreshAccessTokenResponse struct {
	AccessToken          string `json:"access_token"`
	AccessTokenExpiresAt uint64 `json:"access_token_expires_at"`
}
