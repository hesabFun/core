package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/**
 * @api {post} /v1/companies/:id/categories Insert New Category
 * @apiName InsertNewCategory
 * @apiGroup Categories
 * @apiVersion 0.1.0
 *
 * @apiUse jwt
 *
 * @apiParam (Request body) {Number} [parent_id] Parent ID.
 * @apiParam (Request body) {String} name Category name.
 * @apiParam (Request body) {Number} order Order.
 */
func insertNewCategory(c *gin.Context) {
	var category Categorise

	if err := c.ShouldBindWith(&category, binding.JSON); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	category.CompanyId = c.MustGet("company_id").(uint)
	_, err := MySql.InsertInto("product_categories").Values(category).Exec()

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "successful"})
	return
}

type Categorise struct {
	Id        uint   `db:"id"`
	ParentId  uint   `db:"parent_id" json:"parent_id" binding:"lte=2147483648"`
	CompanyId uint   `db:"company_id"`
	Name      string `db:"name" json:"name" binding:"required,gte=0,lte=64"`
	Order     uint   `db:"order" json:"order" binding:"lte=2147483648"`
}
