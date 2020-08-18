package controllers

import(
	"server/src/models"
)

type CompanyController struct{
	companyModel models.CompanyModel
}

func NewCompanyController() CompanyController {
	return CompanyController{}
}