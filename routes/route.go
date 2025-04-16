package routes

import (
	"pos/controller"
	"pos/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	product := r.Group("/products")
	activity := r.Group("/activities")

	product.Use(middleware.AuthMiddleware())
	activity.Use(middleware.AuthMiddleware())

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	product.GET("", controller.GetProduct)
	product.POST("", middleware.AdminOnly(), controller.CreateProduct)
	product.PUT("/:id", middleware.AdminOnly(), controller.UpdateProduct)
	product.DELETE("/:id", middleware.AdminOnly(), controller.DeleteProduct)
	activity.POST("", middleware.AdminOnly(), controller.CreateTransaction)

}
