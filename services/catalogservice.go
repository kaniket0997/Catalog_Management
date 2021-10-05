package services

import "github.com/Aniket/restapi/models"

type ICatalogService interface {
	CreateProduct(models.Product) error
	ShowProduct() []models.Product
	BoughtProduct(models.Product) (models.Product,error)
	UpdateProduct(models.Product) (models.Product,error)
	DeleteProduct(product_id int) error
}