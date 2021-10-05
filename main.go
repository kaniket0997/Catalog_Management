package main

import (
	"github.com/Aniket/restapi/api"
	"github.com/Aniket/restapi/implementation"
	"github.com/Aniket/restapi/models"
	//"github.com/Aniket/restapi/ops"
	//"github.com/Aniket/restapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//Init Router
	r:= mux.NewRouter()

	catalogService := &implementation.Inmemoryimplement{ProductArr: []models.Product{}}
	catalogController := api.CatalogController{CatalogService: catalogService}

	//Route handlers endpoints
	r.HandleFunc("/CreateProduct", catalogController.CreateProduct).Methods("POST")
	r.HandleFunc("/ShowProduct", catalogController.ShowProduct).Methods("GET")
	r.HandleFunc("/BoughtProduct/{id}", catalogController.BoughtProduct).Methods("PUT")
	r.HandleFunc("/UpdateProduct/{id}", catalogController.UpdateProduct).Methods("PUT")
	r.HandleFunc("/DeleteProduct/{id}", catalogController.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
