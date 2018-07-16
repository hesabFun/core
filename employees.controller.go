package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/**
 * @api {post} /v1/companies/:id/employees/ Add User To Company
 * @apiName AddUserToCompany
 * @apiGroup Employees
 * @apiVersion 0.1.0
 *
 * @apiUse jwt
 *
 * @apiParam (Request body) {String=none,manager,accountant,headmaster_accountant,technical} type Type of Employee.
 * @apiParam (Request body) {Number} user_id User ID.
 */
func insertNewEmployee(c *gin.Context) {
	var employee Employees

	if err := c.ShouldBindWith(&employee, binding.JSON); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	//check user exist
	err := MySql.Select("id").From("users").
		Where("id", employee.UserId).
		One(&struct{}{})

	if err != nil {
		c.JSON(400, gin.H{"message": "the user_id is wrong!"})
		return
	}

	//check employee existed
	err = MySql.Select("id").From("employees").
		Where("company_id", c.MustGet("company_id").(uint)).
		Where("user_id", employee.UserId).
		One(&struct{}{})

	if err == nil {
		c.JSON(400, gin.H{"message": "employee exist!"})
		return
	}

	//add employee
	employee.CompanyId = c.MustGet("company_id").(uint)
	employee.StatusByCompany = "active"
	employee.StatusByEmployee = "pending"
	_, err = MySql.InsertInto("employees").Values(employee).Exec()

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "successful"})
	return
}

/**
 * @api {get} /v1/companies/:id/employees/ Get Employees List Of Company
 * @apiName GetEmployeesListOfCompany
 * @apiGroup Employees
 * @apiVersion 0.1.0
 *
 * @apiUse jwt
 */
func getAllEmployeesOfCompany(c *gin.Context) {

	var employees []struct {
		Id               uint   `db:"id" json:"id"`
		UserId           uint   `db:"user_id" json:"user_id"`
		StatusByEmployee string `db:"status_by_employee" json:"status_by_employee"`
		StatusByCompany  string `db:"status_by_company" json:"status_by_company"`
		Name             string `db:"name" json:"name"`
	}

	err := MySql.Select(
		"employees.id",
		"employees.user_id",
		"employees.status_by_employee",
		"employees.status_by_company",
		"users.name",
	).From("employees").
		Join("users").On("users.id = employees.user_id").
		Where("employees.company_id", c.MustGet("company_id").(uint)).
		All(&employees)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, employees)
}

/**
 * @api {get} /v1/employees Get Invite To Company List
 * @apiName GetInviteEmployeeToCompanyList
 * @apiGroup Employees
 * @apiVersion 0.1.0
 *
 * @apiUse jwt
 */
func getAllAddMeToEmployeeRequests(c *gin.Context) {

	loginUser := c.MustGet("user").(LoginUser)

	var employees []struct {
		Id               uint   `db:"id" json:"id"`
		StatusByEmployee string `db:"status_by_employee" json:"status_by_employee"`
		StatusByCompany  string `db:"status_by_company" json:"status_by_company"`
		Name             string `db:"name" json:"name"`
	}

	err := MySql.Select(
		"employees.id",
		"employees.status_by_employee",
		"employees.status_by_company",
		"companies.name",
	).From("employees").
		Join("companies").On("companies.id = employees.company_id").
		Where("employees.user_id", loginUser.Id).
		All(&employees)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, employees)
}

/**
 * @api {put} /v1/employees/:id Change Status Employee Request By User
 * @apiName ChangeEmployeeStatus
 * @apiGroup Employees
 * @apiVersion 0.1.0
 *
 * @apiUse jwt
 *
 * @apiParam (Request body) {String=pending,active,block} status Employee status.
 */
func changeEmployeeStatusByUser(c *gin.Context) {

	loginUser := c.MustGet("user").(LoginUser)

	var employeeStatus struct {
		Status string `db:"status" json:"status" binding:"required,oneof=pending active block"`
	}

	if err := c.ShouldBindWith(&employeeStatus, binding.JSON); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err := MySql.Select("id").From("employees").
		Where("id", c.Param("employee_id")).
		Where("user_id", loginUser.Id).
		One(&struct{}{})

	if err != nil {
		c.JSON(400, gin.H{"message": "the employee_id is wrong!"})
		return
	}

	res := MySql.Collection("employees").
		Find(string2Int(c.Param("employee_id")))

	err = res.Update(map[string]string{
		"status_by_employee": employeeStatus.Status,
	})

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

type Employees struct {
	Id               uint   `db:"id" json:"id"`
	CompanyId        uint   `db:"company_id" json:"company_id"`
	UserId           uint   `db:"user_id" json:"user_id" binding:"required"`
	StatusByEmployee string `db:"status_by_employee" json:"status_by_employee"`
	StatusByCompany  string `db:"status_by_company" json:"status_by_company"`
	Type             string `db:"type" json:"type" binding:"required,oneof=none manager accountant headmaster_accountant technical"`
}
