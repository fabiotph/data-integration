package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
	"regexp"
	"server/src/utils"
	"strings"
	"errors"
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

type Companies []Company

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

func (model *CompanyModel) GetAll() (Companies, error){
	db := utils.Connect()
	response := Companies{}
	err := db.Conn.Find(&response).Error
	if err != nil{
		log.Fatalf("Error get all companies")
	}
	defer db.Close()
	return response, err
}

func (model *CompanyModel) Insert (company *Company) (*Company, error) {
	company.Name = strings.ToUpper(company.Name)
	if !model.Prepare(company) {
		return &Company{}, errors.New("Company is not valid.")
	}
	db := utils.Connect()
	err := db.Conn.Create(company).Error
	if err != nil{
		log.Fatalf("Error to persis Company: %v", err)
	}
	return company, err
}

func (model *CompanyModel) GetByNameAndZipCode(company *Company) (Company, error){
	db := utils.Connect()
	response := Company{}
	err := db.Conn.Where("name LIKE ? AND zip_code = ?", "%"+company.Name+"%", company.ZipCode).Take(&response).Error
	if err != nil{
		log.Println("Company not found.")
	}
	defer db.Close()
	return response, err
}