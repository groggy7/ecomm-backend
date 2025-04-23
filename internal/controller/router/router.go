package router

import (
	"ecomm/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(ph *controller.ProductHandler) *gin.Engine {
	engine := gin.Default()

	engine.GET("/products", ph.GetAllProducts)
	engine.GET("/products/:id", ph.GetProductByID)
	engine.POST("/products", ph.CreateProduct)
	engine.PUT("/products/:id", ph.UpdateProduct)
	engine.DELETE("/products/:id", ph.DeleteProduct)

	return engine
}
