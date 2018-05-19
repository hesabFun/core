package main

import "github.com/gin-gonic/gin"

func getAllCompanies(c *gin.Context) {
	var companies []Companies

	err := MySql.Select(
		"companies.id as id",
		"companies.name as name",
		"companies.status as status",
		"companies.created_at as created_at",
		"companies.deleted_at as deleted_at",
	).From("companies").
		Where("employees.user_id LIKE ?", loginUser.Id).
		Join("employees").On("employees.company_id = companies.id").
		All(&companies)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": companies})
	return
}

type Companies struct {
	Id        uint   `db:"id" json:"id"`
	Name      string `db:"name" json:"name" binding:"required,gte=0,lte=64"`
	Status    string `db:"status" json:"status" binding:"required,oneof=active pending block"`
	CreatedAt string `db:"created_at" json:"created_at"`
	DeletedAt string `db:"deleted_at" json:"deleted_at"`
}
