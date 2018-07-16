package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"os"
	"time"
)

func registerNewUser(c *gin.Context) {
	//validate data
	var form struct {
		//Email    string `db:"email" json:"email" binding:"email"`
		Mobile   string `json:"mobile" binding:"required"`
		Name     string `json:"name" binding:"lte=30"`
		Password string `json:"password" binding:"required,gte=7,lte=130"`
	}

	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	//if form.Email == "" && form.Mobile == 0 {
	//	c.JSON(400, gin.H{"message": "email or mobile can't found!"})
	//	return
	//}

	//check user email or number existed
	err := MySql.Collection("users").Find().Where("mobile", form.Mobile).One(&struct{}{})
	if err == nil {
		c.JSON(400, gin.H{"message": "the phone number existed!"})
		return
	}

	//insert to db
	user := User{
		Name:      form.Name,
		Mobile:    form.Mobile,
		Password:  GetMD5Hash(form.Password),
		Status:    "pending",
		Type:      "user",
		SmsToken:  randomInt(1000, 9999),
		DeletedAt: time.Time{},
		CreatedAt: time.Now(),
	}

	newUser, err := MySql.InsertInto("users").Values(user).Exec()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	//send token
	newUserId, _ := newUser.LastInsertId()
	user.ID = uint(newUserId)

	token, err := generateToken(user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if os.Getenv("GIN_MODE") == "debug" {
		c.JSON(201, gin.H{
			"sms_token": user.SmsToken,
			"token":     token,
		})
		return
	}

	c.JSON(201, gin.H{
		"token": token,
	})
	return
}

func verifyUserBySms(c *gin.Context) {
	var form struct {
		SmsToken int    `json:"sms_token" binding:"required,gte=999,lte=10000"`
		Mobile   string `json:"mobile" binding:"required"`
	}
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	var user User
	err := MySql.Select("*").From("users").
		Where("mobile", form.Mobile).
		Where("sms_token", form.SmsToken).
		One(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "token is wrong!"})
		return
	}

	user.SmsToken = 0
	user.Status = "active"
	err = MySql.Collection("users").Find().Where("mobile", form.Mobile).Update(user)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}
