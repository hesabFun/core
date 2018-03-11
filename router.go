package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/v1/auth/login", loginController)
	router.GET("/v1/auth/user", jwtAuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": loginUser,
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	return router
}
