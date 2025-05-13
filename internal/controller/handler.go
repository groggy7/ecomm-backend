package controller

import (
	"context"
	"ecomm/internal/adapters"
	"ecomm/internal/controller/auth"
	"ecomm/internal/domain"
	"ecomm/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	client     proto.ApiServiceClient
	jwtManager *auth.JWTManager
}

func NewHandler(client proto.ApiServiceClient) *Handler {
	jwtManager, err := auth.NewTokenGenerator()
	if err != nil {
		panic(err)
	}

	return &Handler{
		client:     client,
		jwtManager: jwtManager,
	}
}

func (ph *Handler) CreateProduct(ctx *gin.Context) {
	var request domain.CreateProductRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createRequest := adapters.ToProtoCreateProductRequest(&request)
	createdProduct, err := ph.client.CreateProduct(context.Background(), createRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, createdProduct)
}

func (ph *Handler) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := ph.client.GetProductByID(context.Background(), &proto.GetProductByIDRequest{Id: id})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		ctx.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(200, product)
}

func (ph *Handler) ListProducts(ctx *gin.Context) {
	products, err := ph.client.ListProducts(context.Background(), &proto.ListProductsRequest{})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, products)
}

func (ph *Handler) UpdateProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	if productID == "" {
		ctx.JSON(400, gin.H{"error": "Product ID is required"})
		return
	}

	var request *domain.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	request.ID = productID
	updateRequest := adapters.ToProtoUpdateProductRequest(request)
	_, err := ph.client.UpdateProduct(context.Background(), updateRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Product updated successfully"})
}

func (ph *Handler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := ph.client.DeleteProduct(context.Background(), &proto.DeleteProductRequest{Id: id})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Product deleted successfully"})
}

func (ph *Handler) CreateOrder(ctx *gin.Context) {
	claims, err := ph.jwtManager.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "claims not found"})
		return
	}

	var request domain.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.UserID = claims.ID
	createRequest := adapters.ToProtoCreateOrderRequest(&request)
	order, err := ph.client.CreateOrder(context.Background(), createRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

func (ph *Handler) GetOrder(ctx *gin.Context) {
	claims, err := ph.jwtManager.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "claims not found"})
		return
	}

	order, err := ph.client.GetOrder(context.Background(), &proto.GetOrderRequest{UserId: claims.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if order == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(200, order)
}

func (ph *Handler) ListOrders(ctx *gin.Context) {
	orders, err := ph.client.ListOrders(context.Background(), &proto.ListOrdersRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, orders)
}

func (ph *Handler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := ph.client.DeleteOrder(context.Background(), &proto.DeleteOrderRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (ph *Handler) CreateUser(ctx *gin.Context) {
	var request domain.CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createRequest := adapters.ToProtoCreateUserRequest(&request)
	user, err := ph.client.CreateUser(context.Background(), createRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (ph *Handler) ListUsers(ctx *gin.Context) {
	users, err := ph.client.ListUsers(context.Background(), &proto.ListUsersRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (ph *Handler) UpdateUser(ctx *gin.Context) {
	claims, err := ph.jwtManager.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "claims not found"})
		return
	}

	var request domain.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.ID = claims.ID
	updateRequest := adapters.ToProtoUpdateUserRequest(&request)
	_, err = ph.client.UpdateUser(context.Background(), updateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (ph *Handler) DeleteUser(ctx *gin.Context) {
	claims, err := ph.jwtManager.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "claims not found"})
		return
	}

	_, err = ph.client.DeleteUser(context.Background(), &proto.DeleteUserRequest{
		UserId:    claims.ID,
		SessionId: claims.RegisteredClaims.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (ph *Handler) Login(ctx *gin.Context) {
	var request domain.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginRequest := adapters.ToProtoLoginUserRequest(&request)
	loginResponse, err := ph.client.Login(context.Background(), loginRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": loginResponse})
}

func (ph *Handler) Logout(ctx *gin.Context) {
	claims, err := ph.jwtManager.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "claims not found"})
		return
	}

	_, err = ph.client.Logout(context.Background(), &proto.LogoutRequest{SessionId: claims.RegisteredClaims.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (ph *Handler) RefreshAccessToken(ctx *gin.Context) {
	var request domain.RefreshAccessTokenRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := ph.jwtManager.ValidateToken(request.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "claims not found"})
		return
	}

	request.SessionID = claims.RegisteredClaims.ID
	refreshRequest := adapters.ToProtoRefreshTokenRequest(&request)
	token, err := ph.client.RefreshToken(context.Background(), refreshRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (ph *Handler) RevokeSession(ctx *gin.Context) {
	claims, err := ph.jwtManager.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "claims not found"})
		return
	}

	_, err = ph.client.RevokeSession(context.Background(), &proto.RevokeSessionRequest{SessionId: claims.RegisteredClaims.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Session revoked successfully"})
}
