package main

import(
	"server/src/controllers"
	"server/src/utils"
	"server/src/models"
	"server/src/routes"
	"net/http"
	"log"
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
	log.Println("Server listen on :3000")
	http.ListenAndServe(":3000", router)

}