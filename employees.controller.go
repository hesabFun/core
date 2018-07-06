package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func insertNewEmploye(c *gin.Context) {
	var employee Employees

	if err := c.ShouldBindWith(&employee, binding.JSON); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	employee.CompanyId = c.MustGet("company_id").(uint)
	_, err := MySql.InsertInto("transactions").Values(employee).Exec()

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "successful"})
	return
}

type Employees struct {
	Id        uint   `db:"id" json:"id"`
	CompanyId uint   `db:"company_id" json:"company_id"`
	UserId    uint   `db:"user_id" json:"user_id"`
	Status    string `db:"status" json:"status" binding:"required,oneof=pending active block"`
	Type      string `db:"type" json:"type" binding:"required,oneof=none manager accountant headmaster_accountant technical"`
	Title     string `db:"title" json:"title" binding:"lte=64"`
}
