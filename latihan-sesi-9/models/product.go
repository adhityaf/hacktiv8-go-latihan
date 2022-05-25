package models

import "fmt"

type Product struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Stock int    `json:"stock"`
	Price int    `json:"price"`
}

var Products []Product

func GetProducts() *[]Product {
	return &Products
}

func AddNewProduct(product *Product) {
	Products = append(Products, *product)
}

func GetProductByBrand(brand string) (*[]Product, error) {
	var p []Product
	isExist := false
	for _, product := range Products {
		if product.Brand == brand {
			p = append(p, product) 
			isExist = true
		}
	}

	if !isExist {
		return nil, fmt.Errorf("no product with brand named %s", brand)
	}
	return &p, nil
}
