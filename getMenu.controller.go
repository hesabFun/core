package main

import "github.com/gin-gonic/gin"

func getMenu(c *gin.Context) {

	c.JSON(200, gin.H{
		"menu": "items",
	})
	return
}
