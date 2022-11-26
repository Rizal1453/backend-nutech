package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nutech/dto"
	authdto "nutech/dto/auth"
	"nutech/models"
	"nutech/pkg/bcrypt"
	jwttoken "nutech/pkg/jwt"
	"nutech/repository"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	
)

type handlerAuth struct {
	AuthRepository repository.AuthRepository
}

func HandlerAuth(AuthRepository repository.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	user := models.User{
		Name : request.Name,
		Email:    request.Email,
		Password: password,
	
	}
	data, err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccesResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}
func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.LoginRequest)

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Bad Request"}
		json.NewEncoder(w).Encode(response)
		return
	}
	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}
	// check email
	User, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "email error"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// check password
	isValid := bcrypt.CheckPasswordHash(request.Password, User.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "password salah"}
		json.NewEncoder(w).Encode(response)
		return
	}
	claims := jwt.MapClaims{}
	claims["id"] = User.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwttoken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}
	loginResponse := authdto.LoginResponse{
		ID:       User.ID,
		Email:    User.Email,
		Name: User.Name,
		Token :token,
	}
	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccesResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	user, err := h.AuthRepository.CheckAuth(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:       user.ID,
		Name: user.Name,
		Email:    user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccesResult{Code: http.StatusOK, Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}
