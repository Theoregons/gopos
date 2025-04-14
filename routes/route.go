package routes

import (
	"pos/controller"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	product := r.Group("/products")
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	product.GET("", controller.GetProduct)
	product.POST("", controller.CreateProduct)
	product.PUT("/:id", controller.UpdateProduct)
	product.DELETE("/:id", controller.DeleteProduct)
}
