package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin,Content-Length,Content-Type,Authorization"}
	router.Use(cors.New(corsConfig))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	router.POST("/v1/auth/login", loginController)

	v1 := router.Group("/v1", jwtAuthMiddleware(), rbacCompanyMiddleware())

	authRoute := v1.Group("/auth")
	authRoute.GET("/profile", profileController)

	companies := v1.Group("/companies")
	companies.GET("/")
	companies.GET("/:id")
	companies.GET("/:id/menu", getMenu)

	companies.GET("/:id/products")
	//companies.POST("/:id/products", insertNewProduct)

	companies.POST("/:id/categories", insertNewCategory)

	return router
}
