package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
)

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

func (company *Company) BeforeCreate(scope *gorm.Scope) error{
	err := scope.SetColumn("Id", uuid.NewV4().String())
	if err != nil{
		log.Fatalf("Error during Company Object creating: %v", err)
	}
	return err
}