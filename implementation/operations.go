package implementation

import (
	//"encoding/json"
	"errors"
	"github.com/Aniket/restapi/models"
	//"github.com/gorilla/mux"
	//"strconv"
)

type Inmemoryimplement struct {
	ProductArr []models.Product
}

func (h *Inmemoryimplement)  CreateProduct(product models.Product) error{
	h.ProductArr= append(h.ProductArr, product)

	return nil
}

func (h *Inmemoryimplement) ShowProduct () []models.Product{
	return h.ProductArr
}

func (h *Inmemoryimplement) BoughtProduct (p models.Product) (models.Product, error){
	u:= models.Product{}
	for i, v := range h.ProductArr {
		if v.ID == p.ID {
			if v.Quantity >= p.Quantity {
				v.Quantity -= p.Quantity
				h.ProductArr[i] = v
				//idx:= i
				//json.NewEncoder(w).Encode(h.ProductArr[idx])
				return v,nil
			} else {
				return u,errors.New("product quantity exceeds inventory")
			}
		}
	}
	return u,errors.New("product with given id doesn't exist")
}

func (h *Inmemoryimplement) DeleteProduct (product_id int) error {
	for index, item := range h.ProductArr{
		if item.ID== product_id{
		h.ProductArr = append(h.ProductArr[:index], h.ProductArr[index+1:]...)
		break
		}
	}
	return nil
}

func (h *Inmemoryimplement) UpdateProduct (p models.Product) (models.Product, error){
	u:= models.Product{}
	for i, v := range h.ProductArr {
		if v.ID == p.ID {
				v.Quantity += p.Quantity
				h.ProductArr[i] = v

				return v,nil
		}
	}
	return u,errors.New("product with given id doesn't exist")
}




