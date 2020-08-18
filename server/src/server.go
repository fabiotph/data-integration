package main

import(
	"server/src/controllers"
	"server/src/utils"
	"server/src/models"
	"server/src/routes"
	"net/http"
)


func initialSetting(){
	controller := controllers.CompanyController{}
	db := utils.Connect()
	if !db.Conn.HasTable(&models.Company{}){
		db.Conn.AutoMigrate(&models.Company{})
		controller.LoadCSV("assets/csv/q1_catalog.csv")
	}
	defer db.Conn.Close()
}

func main() {
	initialSetting()
	router := routes.HandleRoutes()

	http.ListenAndServe(":3000", router)
}