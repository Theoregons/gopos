package routes

import (
	"pos/controller"
	"pos/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	product := r.Group("/products")
	activity := r.Group("/activities")
	report := r.Group("/reports")

	product.Use(middleware.AuthMiddleware())
	activity.Use(middleware.AuthMiddleware())
	report.Use(middleware.AuthMiddleware())

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	product.GET("", controller.GetProduct)
	product.POST("", middleware.AdminOnly(), controller.CreateProduct)
	product.PUT("/:id", middleware.AdminOnly(), controller.UpdateProduct)
	product.DELETE("/:id", middleware.AdminOnly(), controller.DeleteProduct)
	activity.POST("", middleware.AdminOnly(), controller.CreateTransaction)
	activity.GET("", middleware.AdminOnly(), controller.GetTransactions)
	activity.GET("/:id", middleware.AdminOnly(), controller.GetTransactionByID)
	report.GET("", middleware.AdminOnly(), controller.GetReports)

}
