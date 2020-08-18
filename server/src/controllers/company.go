package controllers

import(
	"server/src/models"
	"server/src/utils"
	"net/http"
	"os"
	"log"
	"encoding/csv"
)

type CompanyController struct{
	companyModel models.CompanyModel
}

func NewCompanyController() CompanyController {
	return CompanyController{}
}

func (controller *CompanyController) GetAll(w http.ResponseWriter, r *http.Request){
	result, err := controller.companyModel.GetAll()
	if err != nil {
		utils.JSONResponse(w, result, http.StatusInternalServerError)
	}
	utils.JSONResponse(w, result, http.StatusOK)
	return
}

func (controller *CompanyController) LoadCSV(filename string) bool{
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Couldn't open the csv file: %v", err)
			return false
		}
		reader := csv.NewReader(file)
		reader.Comma = ';'

		lines, err := reader.ReadAll()

		if err != nil{
			log.Fatalf("Couldn't read csv file: %v", err)
			return false
		}
		for i, line := range lines{
			if i != 0{
				company := models.Company{Name: line[0], ZipCode: line[1], Website: ""}
				controller.companyModel.Insert(&company)
			}
		}
		return true
}