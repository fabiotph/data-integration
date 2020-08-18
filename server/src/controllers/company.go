package controllers

import(
	"server/src/models"
	"server/src/utils"
	"net/http"
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