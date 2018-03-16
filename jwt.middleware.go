package main

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/gin-gonic/gin"
	"os"
)

type LoginUser struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

// user login info
var loginUser LoginUser

func jwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("JWT_PUBLIC_KEY")))
		jwt, err := jws.ParseJWTFromRequest(c.Request)
		if err != nil {
			respondWithError(401, err.Error(), c)
			return
		}

		if err = jwt.Validate(rsaPublic, crypto.SigningMethodRS256); err != nil {
			respondWithError(401, err.Error(), c)
			return
		}

		temp := jwt.Claims().Get("user").(map[string]interface{})

		loginUser.Id, _ = jwt.Claims().Subject()
		loginUser.Name = temp["name"].(string)
		loginUser.Status = temp["status"].(string)
		loginUser.Type = temp["type"].(string)

		c.Next()
	}
}
