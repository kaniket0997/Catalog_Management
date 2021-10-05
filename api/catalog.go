package api

import (
	"encoding/json"
	"github.com/Aniket/restapi/models"
	"github.com/Aniket/restapi/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CatalogController struct {
	CatalogService services.ICatalogService
}

func (h* CatalogController) CreateProduct (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var p models.Product
	_ = json.NewDecoder(r.Body).Decode(&p)
	err:= h.CatalogService.CreateProduct(p)
	if err!=nil {
		json.NewEncoder(w).Encode("Encountered while adding product.")
		return
	}

	json.NewEncoder(w).Encode("Product added successfully")
}

func (h* CatalogController) ShowProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	product:= h.CatalogService.ShowProduct()
	json.NewEncoder(w).Encode(product)
}

func (h *CatalogController) BoughtProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var p models.Product
	json.NewDecoder(r.Body).Decode(&p)
	prod, err:= h.CatalogService.BoughtProduct(p)
	if err==nil {
		json.NewEncoder(w).Encode("Product successfully bought!!")
		json.NewEncoder(w).Encode(prod)
		return
	}
	json.NewEncoder(w).Encode("product with given id doesn't exist")
	//w.WriteHeader(404)
	//w.Write([]byte("Product with given id doesn't exist !!"))
}


func (h *CatalogController) DeleteProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)["id"]
	id,_:= strconv.Atoi(params)
	h.CatalogService.DeleteProduct(id)
	json.NewEncoder(w).Encode("Deletion of product was successfull")
}


func (h *CatalogController) UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var p models.Product
	json.NewDecoder(r.Body).Decode(&p)
	prod, err:= h.CatalogService.UpdateProduct(p)
	if err==nil {
		json.NewEncoder(w).Encode("Product successfully bought!!")
		json.NewEncoder(w).Encode(prod)
		return
	}
	json.NewEncoder(w).Encode("product with given id doesn't exist")

}





