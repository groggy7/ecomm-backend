package controller

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(ph *ProductHandler) *gin.Engine {
	engine := gin.Default()

	engine.POST("/products", ph.CreateProduct)
	engine.GET("/products", ph.GetAllProducts)
	engine.GET("/products/:id", ph.GetProductByID)
	engine.PUT("/products/:id", ph.UpdateProduct)
	engine.DELETE("/products/:id", ph.DeleteProduct)

	engine.POST("/orders", ph.CreateOrder)
	engine.GET("/orders", ph.GetAllOrders)
	engine.GET("/orders/:id", ph.GetOrderByID)
	engine.DELETE("/orders/:id", ph.DeleteOrder)

	engine.POST("/users", ph.CreateUser)
	engine.GET("/users", ph.ListUsers)
	engine.PUT("/users", ph.UpdateUser)
	engine.DELETE("/users/:id", ph.DeleteUser)

	engine.POST("/login", ph.Login)
	engine.POST("/logout", ph.Logout)
	engine.POST("/sessions/refresh", ph.RefreshAccessToken)
	engine.POST("/sessions/revoke/:id", ph.RevokeSession)
	return engine
}
