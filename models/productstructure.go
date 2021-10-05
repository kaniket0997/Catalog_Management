package models

type Product struct {
	ID int `json:"id"`
	Name string `json:"product_name"`
	Description string `json:"desc"`
	Price  int `json:"price"`
	Quantity int `json:"quantity"`
}
