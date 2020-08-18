package routes

import (
	"github.com/gorilla/mux"
	"server/src/controllers"
	"net/http"
)

func HandleRoutes() *mux.Router {
	controller := controllers.CompanyController{}
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/company/all", controller.GetAll).Methods(http.MethodGet)

	return router
}