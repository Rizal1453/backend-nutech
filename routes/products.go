package routes

import (
	"nutech/handlers"
	"nutech/pkg/middleware"
	"nutech/pkg/mysql"
	"nutech/repository"

	"github.com/gorilla/mux"
)

func ProductRoute(r *mux.Router){
	productRepository := repository.RepositoryProduct(mysql.DB)
	h:= handlers.HandlerProduct(productRepository)

	r.HandleFunc("/create/products",middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")
	r.HandleFunc("/products",h.FindProduct).Methods("GET")
	r.HandleFunc("/product/{id}",h.GetProduct).Methods("GET")
	r.HandleFunc("/update/{id}",middleware.Auth(middleware.UploadFile(h.UpdateProduct))).Methods("PATCH")
	r.HandleFunc("/delete/{id}",middleware.Auth(h.DeleteProduct) ).Methods("DELETE")
	
}