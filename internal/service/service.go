package service

import (
	"context"
	"ecomm/internal/adapters"
	"ecomm/internal/controller/auth"
	"ecomm/internal/domain"
	"ecomm/proto"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	repo       domain.Repository
	jwtManager *auth.JWTManager
	proto.UnimplementedApiServiceServer
}

func NewService(repo domain.Repository) proto.ApiServiceServer {
	jwtManager, err := auth.NewTokenGenerator()
	if err != nil {
		panic(err)
	}
	return &service{
		repo:       repo,
		jwtManager: jwtManager,
	}
}

func (s *service) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	product := &domain.Product{
		Name:            req.Name,
		Image:           req.Image,
		Category:        req.Category,
		Description:     req.Description,
		Rating:          int(req.Rating),
		NumberOfReviews: int(req.NumberOfReviews),
		Price:           float64(req.Price),
		CountInStock:    int(req.CountInStock),
	}

	product, err := s.repo.CreateProduct(product)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}
	return &proto.CreateProductResponse{
		Product: adapters.ToProtoProduct(*product),
	}, nil
}

func (s *service) GetProductByID(ctx context.Context, req *proto.GetProductByIDRequest) (*proto.GetProductByIDResponse, error) {
	product, err := s.repo.GetProductByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get product: %v", err)
	}
	return &proto.GetProductByIDResponse{
		Product: adapters.ToProtoProduct(*product),
	}, nil
}

func (s *service) ListProducts(ctx context.Context, req *proto.ListProductsRequest) (*proto.ListProductsResponse, error) {
	products, err := s.repo.ListProducts()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list products: %v", err)
	}

	return &proto.ListProductsResponse{
		Products: adapters.ToProtoProducts(products),
	}, nil
}

func (s *service) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) (*proto.UpdateProductResponse, error) {
	product, err := s.repo.GetProductByID(req.Id)
	if err != nil {
		return nil, err
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
		product.Rating = int(req.Rating)
	}
	if req.NumberOfReviews != 0 {
		product.NumberOfReviews = int(req.NumberOfReviews)
	}
	if req.Price != 0 {
		product.Price = req.Price
	}
	if req.CountInStock != 0 {
		product.CountInStock = int(req.CountInStock)
	}

	if err := s.repo.UpdateProduct(product); err != nil {
		return nil, err
	}

	return &proto.UpdateProductResponse{}, nil
}

func (s *service) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.DeleteProductResponse, error) {
	if err := s.repo.DeleteProduct(req.Id); err != nil {
		return nil, err
	}
	return &proto.DeleteProductResponse{
		Id: req.Id,
	}, nil
}

func (s *service) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	for _, item := range req.OrderItems {
		_, err := s.repo.GetProductByID(item.ProductId)
		if err != nil {
			return nil, err
		}
	}

	orderItems := make([]*domain.OrderItem, len(req.OrderItems))
	for i, item := range req.OrderItems {
		orderItems[i] = &domain.OrderItem{
			ID:        item.Id,
			OrderID:   item.OrderId,
			ProductID: item.ProductId,
			Name:      item.Name,
			Quantity:  int(item.Quantity),
			Image:     item.Image,
			Price:     item.Price,
		}
	}

	order := &domain.Order{
		PaymentMethod: req.PaymentMethod,
		TaxPrice:      req.TaxPrice,
		ShippingPrice: req.ShippingPrice,
		TotalPrice:    req.TotalPrice,
		OrderItems:    orderItems,
		UserID:        req.UserId,
	}

	order, err := s.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return &proto.CreateOrderResponse{
		Order: adapters.ToProtoOrder(*order),
	}, nil
}

func (s *service) GetOrder(ctx context.Context, req *proto.GetOrderRequest) (*proto.GetOrderResponse, error) {
	order, err := s.repo.GetOrder(req.UserId)
	if err != nil {
		return nil, err
	}

	return &proto.GetOrderResponse{
		Order: adapters.ToProtoOrder(*order),
	}, nil
}

func (s *service) ListOrders(ctx context.Context, req *proto.ListOrdersRequest) (*proto.ListOrdersResponse, error) {
	orders, err := s.repo.ListOrders()
	if err != nil {
		return nil, err
	}

	return &proto.ListOrdersResponse{
		Orders: adapters.ToProtoOrders(orders),
	}, nil
}

func (s *service) DeleteOrder(ctx context.Context, req *proto.DeleteOrderRequest) (*proto.DeleteOrderResponse, error) {
	if err := s.repo.DeleteOrder(req.Id); err != nil {
		return nil, err
	}
	return &proto.DeleteOrderResponse{
		Id: req.Id,
	}, nil
}

func (s *service) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
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

	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &proto.CreateUserResponse{
		Id:      createdUser.ID,
		Name:    createdUser.Name,
		Email:   createdUser.Email,
		IsAdmin: createdUser.IsAdmin,
	}, nil
}

func (s *service) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := s.repo.GetUser(req.Email)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserResponse{
		User: adapters.ToProtoUser(*user),
	}, nil
}

func (s *service) ListUsers(ctx context.Context, req *proto.ListUsersRequest) (*proto.ListUsersResponse, error) {
	users, err := s.repo.ListUsers()
	if err != nil {
		return nil, err
	}

	return &proto.ListUsersResponse{
		Users: adapters.ToProtoUsers(users),
	}, nil
}

func (s *service) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	user, err := s.repo.GetUser(req.Email)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.IsAdmin != user.IsAdmin {
		user.IsAdmin = req.IsAdmin
	}

	user.UpdatedAt = uint64(time.Now().Unix())
	if err := s.repo.UpdateUser(user); err != nil {
		return nil, err
	}

	return &proto.UpdateUserResponse{
		User: adapters.ToProtoUser(*user),
	}, nil
}

func (s *service) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	if err := s.repo.DeleteUser(req.UserId); err != nil {
		return nil, err
	}

	if err := s.repo.DeleteSession(req.SessionId); err != nil {
		return nil, err
	}

	return &proto.DeleteUserResponse{}, nil
}

func (s *service) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := s.repo.GetUser(req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	UUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	tokenID := UUID.String()
	accessToken, _, err := s.jwtManager.GenerateToken(user.Email, user.ID, tokenID, user.IsAdmin, time.Now().Add(3*time.Hour))
	if err != nil {
		return nil, err
	}

	refreshToken, refreshClaims, err := s.jwtManager.GenerateToken(user.Email, user.ID, tokenID, user.IsAdmin, time.Now().Add(3*24*time.Hour))
	if err != nil {
		return nil, err
	}

	if err := s.repo.CreateSession(&domain.Session{
		ID:           refreshClaims.RegisteredClaims.ID,
		Email:        user.Email,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    uint64(refreshClaims.RegisteredClaims.ExpiresAt.Unix()),
	}); err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		SessionId:    refreshClaims.RegisteredClaims.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *service) Logout(ctx context.Context, req *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	session, err := s.repo.GetSession(req.SessionId)
	if err != nil {
		return nil, err
	}

	if session.IsRevoked {
		return nil, fmt.Errorf("invalid session")
	}

	if err := s.repo.RevokeSession(req.SessionId); err != nil {
		return nil, err
	}

	return &proto.LogoutResponse{}, nil
}

func (s *service) RefreshToken(ctx context.Context, req *proto.RefreshAccessTokenRequest) (*proto.RefreshAccessTokenResponse, error) {
	session, err := s.repo.GetSession(req.SessionId)
	if err != nil {
		return nil, err
	}

	if session.ExpiresAt > uint64(time.Now().Unix()) {
		if err := s.repo.RevokeSession(req.SessionId); err != nil {
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

	user, err := s.repo.GetUser(session.Email)
	if err != nil {
		return nil, err
	}

	token, _, err := s.jwtManager.GenerateToken(user.Email, user.ID, session.ID, user.IsAdmin, time.Now().Add(3*time.Hour))
	if err != nil {
		return nil, err
	}

	return &proto.RefreshAccessTokenResponse{
		AccessToken: token,
	}, nil
}

func (s *service) RevokeSession(ctx context.Context, req *proto.RevokeSessionRequest) (*proto.RevokeSessionResponse, error) {
	if err := s.repo.RevokeSession(req.SessionId); err != nil {
		return nil, err
	}
	return &proto.RevokeSessionResponse{}, nil
}
