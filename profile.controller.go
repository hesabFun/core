package main

import "github.com/gin-gonic/gin"

func profileController(c *gin.Context) {
	loginUser := c.MustGet("user").(LoginUser)
	c.JSON(200, gin.H{
		"id":        loginUser.Id,
		"username":  "erfun",
		"firstname": "ErFUN",
		"surname":   "KH",
		"email":     "erfun@email.com",
	})
}
