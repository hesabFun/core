package main

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

type LoginUser struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

func jwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		keyBytes, err := ioutil.ReadFile(os.Getenv("JWT_RSA_PUBLIC_KEY_PATH"))
		if err != nil {
			respondWithError(401, err.Error(), c)
			return
		}

		rsaPublic, err := crypto.ParseRSAPublicKeyFromPEM([]byte(keyBytes))
		if err != nil {
			respondWithError(401, err.Error(), c)
			return
		}

		jwt, err := jws.ParseJWTFromRequest(c.Request)
		if err != nil {
			respondWithError(401, err.Error(), c)
			return
		}

		if err = jwt.Validate(rsaPublic, crypto.SigningMethodRS256); err != nil {
			respondWithError(401, err.Error(), c)
			return
		}

		// user login info
		var loginUser LoginUser

		temp := jwt.Claims().Get("user").(map[string]interface{})

		loginUser.Id, _ = jwt.Claims().Subject()
		loginUser.Name = temp["name"].(string)
		loginUser.Status = temp["status"].(string)
		loginUser.Type = temp["type"].(string)

		c.Set("user", loginUser)

		c.Next()
	}
}
