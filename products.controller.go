package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/**
 * @api {post} /v1/companies/:id/products Insert New Product
 * @apiName InsertNewProduct
 * @apiGroup Products
 * @apiVersion 0.1.0
 *
 * @apiUse jwt
 *
 * @apiParam (Request body) {Number} category_id Category ID.
 * @apiParam (Request body) {String} name Product name.
 * @apiParam (Request body) {String} [description] Product description.
 * @apiParam (Request body) {Number} [price] Product price (0 is free).
 */
func insertNewProduct(c *gin.Context) {
	var product Products
	companyId := c.MustGet("company_id").(uint)

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

/**
 * @api {get} /v1/companies/:id/products Get Products List
 * @apiName GetProductsList
 * @apiGroup Products
 * @apiVersion 0.1.0
 *
 * @apiUse jwt
 */
func getAllProducts(c *gin.Context) {
	var products []Products

	err := MySql.Collection("products").Find().Where("company_id LIKE ?", c.MustGet("company_id").(uint)).All(&products)
	//productsCollection
	//productsCollection.
	//err := MySql.Select("*").From("products").
	//	Where("company_id LIKE ?", companyId).
	//	All(&products)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": products})
}

type Products struct {
	Id          uint   `db:"id" json:"id"`
	CompanyId   uint   `db:"company_id" json:"company_id"`
	CategoryId  uint   `db:"category_id" json:"category_id" binding:"required,gte=0,lte=2147483648"`
	Name        string `db:"name" json:"name" binding:"required,gte=0,lte=64"`
	Description string `db:"description" json:"description"`
	Price       uint   `db:"price" json:"price" binding:"lte=2147483648"`
}
