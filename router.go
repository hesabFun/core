package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	return router
}
