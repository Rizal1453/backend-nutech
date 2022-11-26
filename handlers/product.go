package handlers

import (
	"encoding/json"
	"fmt"

	"context"
	"net/http"
	"nutech/dto"
	productdto "nutech/dto/product"
	"nutech/models"
	"nutech/repository"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	// "github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type handler struct {
	ProductRepository repository.ProductRepository
}
func HandlerProduct(ProductRepository repository.ProductRepository) *handler {
	return &handler{ProductRepository}
}
func (h *handler)FindProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	products,err := h.ProductRepository.FindProduct()
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	for i, p := range products {
		imagePath := os.Getenv("PATH_FILE") + p.Image
		products[i].Image = imagePath
	}
	

	w.WriteHeader(http.StatusOK)
	response := dto.SuccesResult{ Code: http.StatusOK,Data: products}
	json.NewEncoder(w).Encode(response)
}
func (h *handler)CreateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	dataContext := r.Context().Value("dataFile")
	filepath := dataContext.(string) 



	buy, _ := strconv.Atoi(r.FormValue("buy"))
	sale, _ := strconv.Atoi(r.FormValue("sale"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	request := productdto.ProductRequest{
		Name : r.FormValue("name"),
		Buy : buy,
		Sale: sale,
		Qty: qty,
	}
	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "nutech"})

	if err != nil {
		fmt.Println(err.Error())
	}


	product := models.Products{
		Name:   request.Name,
		Image:  resp.SecureURL,
		Buy : request.Buy,
		Sale: request.Sale,
		Qty:    request.Qty,
		
	}
	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	
	w.WriteHeader(http.StatusOK)
	response := dto.SuccesResult{Code: http.StatusOK , Data: product}
	json.NewEncoder(w).Encode(response)
}
func (h *handler)UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filepath := dataContex.(string)             // add this code

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysfood"})

	if err != nil {
		fmt.Println(err.Error())
	}

	
	buy, _ := strconv.Atoi(r.FormValue("buy"))
	sale, _ := strconv.Atoi(r.FormValue("sale"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	request := productdto.UpdateProductRequest{
		Name: r.FormValue("name"),	
		Image:  resp.SecureURL,
		Buy: buy,
		Sale :sale,
		Qty: qty,
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	products, err := h.ProductRepository.GetProduct(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}


	if request.Name != "" {
		products.Name = request.Name

	}


	if request.Image != "" {
		products.Image = request.Image
	}



	if request.Buy != 0 {
		products.Buy = request.Buy
	}

	if request.Sale != 0 {
		products.Sale = request.Sale
	}
	if request.Qty !=0 {
		products.Qty = request.Qty
	}

	data, err := h.ProductRepository.UpdateProduct(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccesResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}


func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ProductRepository.DeleteProduct(product,id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccesResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product models.Products
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	

	w.WriteHeader(http.StatusOK)
	response := dto.SuccesResult{Code: http.StatusOK, Data:product}
	json.NewEncoder(w).Encode(response)
}