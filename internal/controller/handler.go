package controller

import (
	"ecomm/internal/domain"
	"ecomm/internal/usecases"

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

	product := &domain.Product{
		Name:            request.Name,
		Image:           request.Image,
		Category:        request.Category,
		Description:     request.Description,
		Rating:          request.Rating,
		NumberOfReviews: request.NumberOfReviews,
		Price:           request.Price,
		CountInStock:    request.CountInStock,
	}

	createdProduct, err := ph.usecase.CreateProduct(product)
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
	var request domain.UpdateProductRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	product := &domain.Product{
		ID:              id,
		Name:            request.Name,
		Image:           request.Image,
		Category:        request.Category,
		Description:     request.Description,
		Rating:          request.Rating,
		NumberOfReviews: request.NumberOfReviews,
		Price:           request.Price,
		CountInStock:    request.CountInStock,
	}

	err := ph.usecase.UpdateProduct(product)
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
