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

	v1 := router.Group("/v1")

	v1.POST("/auth/login", loginController)
	authRoute := v1.Group("/auth")
	authRoute.Use(jwtAuthMiddleware())
	authRoute.GET("/profile", profileController)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	return router
}
