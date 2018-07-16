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

/**
 * @api {post} /v1/auth/login Request User Login
 * @apiName UserLogin
 * @apiGroup Authorization
 * @apiVersion 0.1.0
 *
 * @apiParam (Request body) {String} mobile User mobile number.
 * @apiParam (Request body) {String} password User password.
 */
func loginController(c *gin.Context) {
	var request struct {
		Mobile   string `json:"mobile" binding:"required,gte=10,lte=12"`
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

	token, err := generateToken(user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func generateToken(user User) (string, error) {

	claims := jws.Claims{
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

	exp, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))

	if err != nil {
		return "", err
	}

	claims.SetExpiration(time.Now().Add(time.Duration(60 * 60 * 24 * time.Duration(exp) * time.Second)))
	claims.SetIssuer(os.Getenv("JWT_ISSUER"))
	claims.SetAudience(os.Getenv("JWT_AUDIENCE"))
	claims.SetIssuedAt(time.Now())
	claims.SetNotBefore(time.Now())
	claims.SetSubject(strconv.FormatUint(uint64(user.ID), 10)) // set user id
	claims.SetJWTID("123")                                     // set token id

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
