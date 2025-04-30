package controller

import (
	"ecomm/internal/controller/auth"

	"github.com/gin-gonic/gin"
)

func NewRouter(ph *ProductHandler) *gin.Engine {
	engine := gin.Default()

	authMiddleware := auth.JWTAuthMiddleware(ph.jwtManager)
	adminMiddleware := auth.JWTAdminMiddleware(ph.jwtManager)

	engine.POST("/products", adminMiddleware, ph.CreateProduct)
	engine.GET("/products", ph.ListProducts)
	engine.GET("/products/:id", ph.GetProductByID)
	engine.PUT("/products/:id", adminMiddleware, ph.UpdateProduct)
	engine.DELETE("/products/:id", adminMiddleware, ph.DeleteProduct)

	engine.POST("/orders", authMiddleware, ph.CreateOrder)
	engine.GET("/orders", authMiddleware, ph.ListOrders)
	engine.GET("/orders/:id", ph.GetOrder)
	engine.DELETE("/orders/:id", authMiddleware, ph.DeleteOrder)

	engine.POST("/users", ph.CreateUser)
	engine.GET("/users", adminMiddleware, ph.ListUsers)
	engine.PUT("/users", authMiddleware, ph.UpdateUser)
	engine.DELETE("/users", adminMiddleware, ph.DeleteUser)

	engine.POST("/login", ph.Login)
	engine.POST("/logout", authMiddleware, ph.Logout)
	engine.POST("/sessions/refresh", authMiddleware, ph.RefreshAccessToken)
	engine.GET("/sessions/revoke", authMiddleware, ph.RevokeSession)
	return engine
}
