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
	companies.GET("", getAllCompanies)
	companies.POST("", insertNewCompany)
	companies.GET("/:companies_id")
	companies.GET("/:companies_id/menu", getMenu)

	companies.GET("/:companies_id/products", getAllProducts)
	companies.POST("/:companies_id/products", insertNewProduct)

	companies.POST("/:companies_id/categories", insertNewCategory)

	companies.GET("/:companies_id/transactions", getAllTransactions)
	companies.POST("/:companies_id/transactions", insertNewTransaction)

	companies.POST("/:companies_id/employees", insertNewEmploye)

	return router
}
