package routes

import (
	"github.com/emanuen/countries-mux/controller"
	"github.com/gorilla/mux"
)

func CountryRoutes(r *mux.Router) {

	r.HandleFunc("/countries", controller.Countries)

}
