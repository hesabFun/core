package main

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"upper.io/db.v3"

	"os"
	"strconv"
	"time"
)

var claims jws.Claims

func loginController(c *gin.Context) {
	var request struct {
		Mobile   string `json:"username" binding:"required,gte=10,lte=12"`
		Password string `json:"password" binding:"required,gte=0,lte=255"`
	}

	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	var user User
	err := MySql.Collection("users").Find(db.Cond{
		"mobile":   request.Mobile,
		"password": GetMD5Hash(request.Password),
	}).Where("status IN ('active', 'pending')").One(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "Mobile or password incorrect"})
		return
	}

	claims = jws.Claims{
		"user": struct {
			Name   string `json:"name"`
			Status string `json:"status"`
			Type   string `json:"type"`
		}{
			Name:   user.Name,
			Status: user.Status,
			Type:   user.Type,
		},
	}

	token, err := generateToken(user.ID)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func generateToken(userId uint) (string, error) {
	exp, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))

	claims.SetExpiration(time.Now().Add(time.Duration(60 * 60 * 24 * time.Duration(exp) * time.Second)))
	claims.SetIssuer(os.Getenv("JWT_ISSUER"))
	claims.SetAudience(os.Getenv("JWT_AUDIENCE"))
	claims.SetIssuedAt(time.Now())
	claims.SetNotBefore(time.Now())
	claims.SetSubject(strconv.FormatUint(uint64(userId), 10)) // set user id
	claims.SetJWTID("123")                                    // set token id

	rsaPrivate, err := crypto.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_PRIVATE_KEY")))

	if err != nil {
		return "", err
	}

	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	token, err := jwt.Serialize(rsaPrivate)

	if err != nil {
		return "", err
	}

	return string(token), nil
}
