package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// role-base access control middleware for companies
// Check https://en.wikipedia.org/wiki/Role-based_access_control for more details
func rbacCompanyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := strings.Split(c.Request.URL.Path, "/")
		dbPath := rbacGetDBPath(path)
		var companyId uint

		if len(path) > 3 {
			u64, _ := strconv.ParseUint(c.Param("companies_id"), 10, 32)
			companyId = uint(u64)
			c.Set("company_id", uint(u64))
		}

		loginUser := c.MustGet("user").(LoginUser)

		if path[2] == "companies" && companyId > 0 {

			// check permission for companies
			res := struct{}{}
			err := MySql.Select("rbac_roles.id as role_id").From("rbac_roles").
				Where("rbac_roles.path LIKE ?", dbPath).
				Where("rbac_roles.method LIKE ?", c.Request.Method).
				Where("rbac_groups.company_id LIKE ?", companyId).
				Where("rbac_group_people.user_id LIKE ?", loginUser.Id).
				Join("rbac_group_roles").On("rbac_group_roles.role_id = rbac_roles.id").
				Join("rbac_group_people").On("rbac_group_people.group_id = rbac_group_roles.group_id").
				Join("rbac_groups").On("rbac_groups.id = rbac_group_roles.group_id").
				One(&res)

			if err != nil {
				respondWithError(403, "permission denied!", c)
				return
			}
		} else {

			//todo: check permission for people
		}

		c.Next()
	}
}

func rbacGetDBPath(path []string) (dbPath string) {
	for _, slice := range path {
		if slice != "" {
			_, err := strconv.Atoi(slice)
			if err == nil {
				slice = ":id"
			}
			dbPath += "/" + slice
		}
	}
	return
}
