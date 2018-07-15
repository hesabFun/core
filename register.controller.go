package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func registerNewUser(c *gin.Context) {
	//validate data
	var form struct {
		//Email    string `db:"email" json:"email" binding:"email"`
		Mobile   string `db:"mobile" json:"mobile" binding:"required"`
		Name     string `db:"name" json:"name" binding:"lte=30"`
		Password string `db:"password" json:"password" binding:"required,gte=0,lte=130"`
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
		Name:     form.Name,
		Mobile:   form.Mobile,
		Password: GetMD5Hash(form.Password),
		Status:   "pending",
		Type:     "user",
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

	c.JSON(201, gin.H{
		"token": token,
	})
	return
}
