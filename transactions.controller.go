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

	if transaction.ProductId.Valid {
		err := MySql.Select("id").From("products").
			Where("id LIKE ?", transaction.ProductId.Int64).
			Where("company_id LIKE ?", companyId).
			One(&struct{}{})
		if err != nil {
			c.JSON(400, gin.H{"message1": err.Error()})
			return
		}
	}

	transaction.CompanyId = companyId
	_, err := MySql.InsertInto("transactions").Columns(
		"company_id",
		"product_id",
		"user_id",
		"title",
		"amount",
		"type",
		"date",
	).Values(
		companyId,
		IfThenElse(transaction.ProductId.Valid, transaction.ProductId.Int64, nil),
		IfThenElse(transaction.UserId.Valid, transaction.UserId.Int64, nil),
		transaction.Title,
		transaction.Amount,
		transaction.Type,
		transaction.Date,
	).Exec()

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "successful"})
	return
}

type Transactions struct {
	Id        uint      `db:"id"`
	CompanyId uint      `db:"company_id"`
	ProductId NullInt64 `db:"product_id" json:"product_id" binding:"lte=2147483648"`
	UserId    NullInt64 `db:"user_id" json:"user_id" binding:"lte=2147483648"`
	Title     string    `db:"title" json:"title" binding:"lte=64"`
	Amount    uint      `db:"amount" json:"amount" binding:"required,gte=0,lte=9223372036854775808"`
	Type      string    `db:"type" json:"type" binding:"required,oneof=input output"`
	Date      string    `db:"date" json:"date" binding:"required"`
}
