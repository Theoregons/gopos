package main

import (
	"pos/config"
	"pos/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	api := r.Group("/api")
	{
		routes.SetupUserRoutes(api)
		// routes.SetupItemRoutes(api)
		// routes.SetupActivityRoutes(api)
		// routes.SetupReportRoutes(api)
	}

	r.Run(":8080")
}
