package main

import (
	"github.com/gin-gonic/gin"
)

func getMenu(c *gin.Context) {
	loginUser := c.MustGet("user").(LoginUser)

	type Menu struct {
		Alias string `db:"alias"`
		Path  string `db:"path"`
	}

	var menu []Menu

	err := MySql.Select("rbac_roles.alias", "rbac_roles.path").From("rbac_roles").
		Join("rbac_group_roles").On("rbac_group_roles.role_id = rbac_roles.id").
		Join("rbac_group_people").On("rbac_group_people.group_id = rbac_group_roles.group_id").
		Join("rbac_groups").On("rbac_groups.id = rbac_group_roles.group_id").
		Where("rbac_roles.menu LIKE ?", "yes").
		Where("rbac_groups.company_id LIKE ?", companyId).
		Where("rbac_group_people.user_id LIKE ?", loginUser.Id).
		All(&menu)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "db error!",
		})
		return
	}

	c.JSON(200, gin.H{
		"menu": menu,
	})
	return
}
