package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func insertNewProduct(c *gin.Context) {
	var product Products

	if err := c.ShouldBindWith(&product, binding.JSON); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err := MySql.Select("id").From("product_categories").
		Where("id LIKE ?", product.CategoryId).
		Where("company_id LIKE ?", companyId).
		One(&struct{}{})
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	product.CompanyId = companyId
	_, err = MySql.InsertInto("products").Values(product).Exec()

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "successful"})
	return
}

type Products struct {
	Id          uint   `db:"id"`
	CompanyId   uint   `db:"company_id"`
	CategoryId  uint   `db:"category_id" json:"category_id" binding:"required,gte=0,lte=2147483648"`
	Name        string `db:"name" json:"name" binding:"required,gte=0,lte=64"`
	Description string `db:"description" json:"description"`
	Price       uint   `db:"price" json:"price" binding:"lte=2147483648"`
}
