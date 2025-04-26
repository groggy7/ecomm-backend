package usecases

import (
	"ecomm/internal/controller/auth"
	"ecomm/internal/domain"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	repo       domain.Repository
	jwtManager *auth.JWTManager
}

func NewUseCase(repo domain.Repository) *UseCase {
	jwtManager, err := auth.NewTokenGenerator()
	if err != nil {
		log.Fatal(err)
	}
	return &UseCase{
		repo:       repo,
		jwtManager: jwtManager,
	}
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

func (u *UseCase) CreateUser(req *domain.CreateUserRequest) (*domain.CreateUserResponse, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
		IsAdmin:  req.IsAdmin,
	}

	createdUser, err := u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &domain.CreateUserResponse{
		ID:      createdUser.ID,
		Name:    createdUser.Name,
		Email:   createdUser.Email,
		IsAdmin: createdUser.IsAdmin,
	}, nil
}

func (u *UseCase) GetUser(email string) (*domain.User, error) {
	return u.repo.GetUser(email)
}

func (u *UseCase) ListUsers() (*domain.ListUserResponse, error) {
	users, err := u.repo.ListUsers()
	if err != nil {
		return nil, err
	}

	resp := &domain.ListUserResponse{
		Users: make([]domain.UserInfo, len(users)),
	}

	for i, user := range users {
		resp.Users[i] = domain.UserInfo{
			ID:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			IsAdmin: user.IsAdmin,
		}
	}

	return resp, nil
}

func (u *UseCase) UpdateUser(req *domain.UpdateUserRequest) error {
	user, err := u.repo.GetUser(req.Email)
	if err != nil {
		return err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.IsAdmin != user.IsAdmin {
		user.IsAdmin = req.IsAdmin
	}

	user.UpdatedAt = uint64(time.Now().Unix())
	return u.repo.UpdateUser(user)
}

func (u *UseCase) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}

func (u *UseCase) Login(req *domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := u.repo.GetUser(req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	accessToken, accessClaims, err := u.jwtManager.GenerateToken(user.Email, user.ID, user.IsAdmin, time.Now().Add(3*time.Hour))
	if err != nil {
		return nil, err
	}

	refreshToken, refreshClaims, err := u.jwtManager.GenerateToken(user.Email, user.ID, user.IsAdmin, time.Now().Add(3*24*time.Hour))
	if err != nil {
		return nil, err
	}

	if err := u.repo.CreateSession(&domain.Session{
		ID:           refreshClaims.RegisteredClaims.ID,
		Email:        user.Email,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    uint64(refreshClaims.RegisteredClaims.ExpiresAt.Unix()),
	}); err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		SessionID:             refreshClaims.RegisteredClaims.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  uint64(accessClaims.RegisteredClaims.ExpiresAt.Unix()),
		RefreshTokenExpiresAt: uint64(refreshClaims.RegisteredClaims.ExpiresAt.Unix()),
	}, nil
}

func (u *UseCase) Logout(req *domain.LogoutRequest) error {
	session, err := u.repo.GetSession(req.SessionID)
	if err != nil {
		return err
	}

	if session.IsRevoked {
		return fmt.Errorf("invalid session")
	}

	if session.Email != req.Email {
		return fmt.Errorf("invalid session")
	}

	if err := u.repo.RevokeSession(req.SessionID); err != nil {
		return err
	}

	return nil
}

func (u *UseCase) RefreshToken(req *domain.RefreshAccessTokenRequest) (*domain.RefreshAccessTokenResponse, error) {
	session, err := u.repo.GetSession(req.SessionID)
	if err != nil {
		return nil, err
	}

	if session.ExpiresAt > uint64(time.Now().Unix()) {
		if err := u.repo.RevokeSession(req.SessionID); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("session expired")
	}

	if session.IsRevoked {
		return nil, fmt.Errorf("session revoked")
	}

	if session.RefreshToken != req.RefreshToken {
		return nil, fmt.Errorf("invalid session")
	}

	user, err := u.repo.GetUser(session.Email)
	if err != nil {
		return nil, err
	}

	token, claims, err := u.jwtManager.GenerateToken(user.Email, user.ID, user.IsAdmin, time.Now().Add(3*time.Hour))
	if err != nil {
		return nil, err
	}

	return &domain.RefreshAccessTokenResponse{
		AccessToken:          token,
		AccessTokenExpiresAt: uint64(claims.RegisteredClaims.ExpiresAt.Unix()),
	}, nil
}

func (u *UseCase) RevokeSession(id string) error {
	return u.repo.RevokeSession(id)
}
