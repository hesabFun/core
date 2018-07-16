package main

import "github.com/gin-gonic/gin"

/**
 * @api {get} /v1/auth/profile Get User information
 * @apiName GetUserProfile
 * @apiGroup Authorization
 * @apiVersion 0.1.0
 *
 * @apiUse jwt
 */

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
