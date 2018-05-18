package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func insertNewTransaction(c *gin.Context) {
	var transaction Transactions

	if err := c.ShouldBindWith(&transaction, binding.JSON); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if transaction.ProductId != nil {
		err := MySql.Select("id").From("products").
			Where("id LIKE ?", transaction.ProductId).
			Where("company_id LIKE ?", companyId).
			One(&struct{}{})
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}
	}

	transaction.CompanyId = companyId
	_, err := MySql.InsertInto("transactions").Values(transaction).Exec()

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "successful"})
	return
}

func getAllTransactions(c *gin.Context) {
	var transactions []Transactions

	err := MySql.Select("*").From("transactions").
		Where("company_id LIKE ?", companyId).
		All(&transactions)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": transactions})
	return
}

type Transactions struct {
	Id        uint   `db:"id" json:"id"`
	CompanyId uint   `db:"company_id" json:"company_id"`
	ProductId *uint  `db:"product_id" json:"product_id"`
	UserId    *uint  `db:"user_id" json:"user_id"`
	Title     string `db:"title" json:"title" binding:"lte=64"`
	Amount    uint   `db:"amount" json:"amount" binding:"required,gte=0,lte=9223372036854775808"`
	Type      string `db:"type" json:"type" binding:"required,oneof=input output"`
	Date      string `db:"date" json:"date" binding:"required"`
}
