package main

import (
	"encoding/json"
	"fmt"
	"latihan-rest-api/middleware"
	"latihan-rest-api/models"
	"net/http"
	"strconv"
)

/*
	Latihan :
	Buatlah sebuah REST API untuk :
		- Register
		  ini cukup dengan menginputkan username dan password.
		  data user bisa lebih dari 1, jadi silahkan gunakan slice.

		- GetAllUsers
		  untuk cek seluruh user. endpoint ini khusus untuk user yg sudah
		  didaftarkan.

		- AddNewProducts
		  yang bisa hit endpoint ini hanyalah user yang sudah di daftarkan.
		  data yang dibutuhkan adalah :
		  	- Nama
			- Brand
			- Stok
			- Price

		- GetProducts
		  tidak ada auth disini. jadi ini adalah API open. akan ngereturn :
		  payload : [
			  {
				  nama 	: "",
				  brand : "",
				  stok 	: 0,
				  price	: 0
			  }
		  ]

		- GetProductByBrand
		  tidak ada auth disini. jadi ini adalah API open.
		  Brand akan di dapat dari query. akan ngereturn :
		  payload : [
			  {
				  nama 	: "",
				  brand : "",
				  stok 	: 0,
				  price	: 0
			  }
		  ]
*/

func main() {
	http.HandleFunc("/register", Register)
	http.HandleFunc("/users", GetAllUsers)

	http.HandleFunc("/postproducts", AddNewProducts)
	http.HandleFunc("/getproducts", GetProducts)

	server := new(http.Server)
	port := ":8080"

	fmt.Println("Server running on port", port)
	server.Addr = port
	server.ListenAndServe()
}

func Register(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var newuser = models.User{
			Username: username,
			Password: password,
		}
		models.AddNewUser(&newuser)
		outputJson(rw, newuser)
		return
	}

	http.Error(rw, "Invalid method", http.StatusBadRequest)
}

func GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	if !middleware.Auth(rw, r){
		return
	}
	outputJson(rw, models.GetUsers())
}

func AddNewProducts(rw http.ResponseWriter, r *http.Request) {
	if !middleware.Auth(rw, r){
		return
	}
	
	rw.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		name := r.FormValue("name")
		brand := r.FormValue("brand")
		stock := r.FormValue("stock")
		price := r.FormValue("price")

		convStock, err := strconv.Atoi(stock)
		if err!=nil{
			http.Error(rw, "Invalid stock", http.StatusBadRequest)
			return
		}

		convPrice, err := strconv.Atoi(price)
		if err!=nil{
			http.Error(rw, "Invalid price", http.StatusBadRequest)
			return
		}

		var newproduct = models.Product{
			Name:  name,
			Brand: brand,
			Stock: convStock,
			Price: convPrice,
		}
		models.AddNewProduct(&newproduct)
		outputJson(rw, newproduct)
		return
	}

	http.Error(rw, "Invalid method", http.StatusBadRequest)
}

func GetProducts(rw http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		query := r.URL.Query()
		brand := query.Get("brand")
		
		if brand == ""{
			outputJson(rw, models.GetProducts())
		}else{
			res, err := models.GetProductByBrand(brand)
			if err != nil {
				outputJson(rw, err.Error())
				return
			}
			outputJson(rw, res)
		}
		return
	}

	http.Error(rw, "Invalid method", http.StatusBadRequest)
}

func outputJson(rw http.ResponseWriter, payload interface{}) {
	response := map[string]interface{}{
		"payload": payload,
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}
