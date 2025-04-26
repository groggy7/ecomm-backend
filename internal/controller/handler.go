package controller

import (
	"ecomm/internal/domain"
	"ecomm/internal/usecases"
	"ecomm/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase *usecases.UseCase
}

func NewProductHandler(usecase *usecases.UseCase) *ProductHandler {
	return &ProductHandler{
		usecase: usecase,
	}
}

func (ph *ProductHandler) CreateProduct(ctx *gin.Context) {
	var request domain.CreateProductRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := ph.usecase.CreateProduct(&request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, createdProduct)
}

func (ph *ProductHandler) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := ph.usecase.GetProductByID(id)
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

func (ph *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := ph.usecase.GetAllProducts()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, products)
}

func (ph *ProductHandler) UpdateProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	if productID == "" {
		ctx.JSON(400, gin.H{"error": "Product ID is required"})
		return
	}

	var request domain.UpdateProductRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	request.ID = productID
	err := ph.usecase.UpdateProduct(&request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Product updated successfully"})
}

func (ph *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ph.usecase.DeleteProduct(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Product deleted successfully"})
}

func (ph *ProductHandler) CreateOrder(ctx *gin.Context) {
	var request domain.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := ph.usecase.CreateOrder(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

func (ph *ProductHandler) GetOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := ph.usecase.GetOrderByID(id)
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

func (ph *ProductHandler) GetAllOrders(ctx *gin.Context) {
	orders, err := ph.usecase.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, orders)
}

func (ph *ProductHandler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ph.usecase.DeleteOrder(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (ph *ProductHandler) CreateUser(ctx *gin.Context) {
	var request domain.CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ph.usecase.CreateUser(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (ph *ProductHandler) ListUsers(ctx *gin.Context) {
	users, err := ph.usecase.ListUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (ph *ProductHandler) UpdateUser(ctx *gin.Context) {
	var request domain.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pkg.ValidateStruct(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := ph.usecase.UpdateUser(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (ph *ProductHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ph.usecase.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (ph *ProductHandler) Login(ctx *gin.Context) {
	var request domain.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ph.usecase.Login(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (ph *ProductHandler) Logout(ctx *gin.Context) {
	var request domain.LogoutRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ph.usecase.Logout(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (ph *ProductHandler) RefreshAccessToken(ctx *gin.Context) {
	var request domain.RefreshAccessTokenRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ph.usecase.RefreshToken(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (ph *ProductHandler) RevokeSession(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Session ID is required"})
		return
	}

	err := ph.usecase.RevokeSession(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Session revoked successfully"})
}
