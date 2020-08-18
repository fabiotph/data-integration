package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
	"regexp"
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

func (model *CompanyModel) Prepare(company *Company) bool{
	if !company.validate(){
		return false
	}
	return true
}

func (model *Company) validate() bool{
	if(!validateName(model.Name) || !validateZipCode(model.ZipCode)) || !validateWebSite(model.Website){
		return false
	}
	return true
}

func validateName(name string) bool{
	isMatch, _ := regexp.MatchString("^([A-ZÁÀÂÃÉÈÊÍÏÓÔÕÖÚÇÑ&']+( [A-ZÁÀÂÃÉÈÊÍÏÓÔÕÖÚÇÑ&'])*)+$", name)
	return isMatch
}

func validateZipCode(zipCode string) bool{
	isMatch, _ := regexp.MatchString("^[0-9]{5}$", zipCode)
	return isMatch
}

func validateWebSite(website string) bool{
	regex := regexp.MustCompile(`^((https?)://[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?((/(([a-z]+[A-Z]*[0-9]*)|([a-z]*[A-Z]+[0-9]*)|([a-z]*[A-Z]*[0-9]+))/?)*|(/?)))?$`)
	isMatch := regex.MatchString(website)
	return isMatch
}
