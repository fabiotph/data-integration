package models

type Company struct{
	Model
	Id string `json:"id" gorm:"type:uuid;primary_key"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	ZipCode string `json:"zip" gorm:"type:varchar(5)"`
	Website string `json:"website" gorm:"type:varchar(255)"`
}

func NewCompany() *Company {
	return &Company{}
}

type CompanyModel struct {}

func NewCompanyModel() *CompanyModel {
	return &CompanyModel{}
}