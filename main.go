package main

import (
	"log"
	"net/http"

	"github.com/emanuen/countries-mux/config"
	"github.com/emanuen/countries-mux/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	r := mux.NewRouter()

	godotenv.Load()

	routes.CountryRoutes(r)

	log.Fatal(http.ListenAndServe(":"+config.Config()["port"], r))

}
