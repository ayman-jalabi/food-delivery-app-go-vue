package handlers

import (
	"encoding/json"
	"main/endpoint"
	"main/models"
	"main/service"
	"net/http"
)

type AuthHandler struct {
	accessConfig  *models.AccessConfig
	refreshConfig *models.RefreshConfig
}

// NewAuthHandler is a constructor for the authHandler struct above
func NewAuthHandler(accessConfig *models.AccessConfig, refreshConfig *models.RefreshConfig) *AuthHandler {
	return &AuthHandler{
		accessConfig:  accessConfig,
		refreshConfig: refreshConfig,
	}
}

// Register is a method for the AuthHandler struct
//func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
//	//this instantiates a kind of object which has the fields that
//	//we stored on the struct we're instantiating
//	//we will use this instance of the struct
//	//to store the email and password obtained from the login request
//	req := new(endpoint.RegistrationRequest)
//
//	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	exists, err := repos.NewUserRepo().CheckEmailExistence(req.Email)
//
//	user, err := repos.NewUserRepo().GetUserByEmail(req.Email)
//	if err != nil {
//		http.Error(w, "invalid credentials", http.StatusUnauthorized)
//		return
//	}
//
//	//we need to compare the password between what the user provided and the one saved on the database
//	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
//		http.Error(w, "invalid credentials", http.StatusUnauthorized)
//		return
//	}
//
//	//here I'm creating an instance of TokenService using the access and refresh config settings
//	//in order to use its methods
//	tokenService := service.NewTokenService(h.accessConfig, h.refreshConfig)
//
//	//as you can see we're using a token service method here to generate an access token
//	accessString, err := tokenService.GenerateAccessToken(user.ID)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	refreshString, err := tokenService.GenerateRefreshToken(user.ID)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	resp := endpoint.LoginResponse{
//		AccessToken:  accessString,
//		RefreshToken: refreshString,
//	}
//
//	w.WriteHeader(http.StatusOK)
//	//encoding resp
//	if err := json.NewEncoder(w).Encode(resp); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}

//func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
//
//	//this stores the email and password provided from within the login request
//	req := new(endpoint.LoginRequest)
//	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	user, err := repos.NewUserRepo().GetUserByEmail(req.Email)
//	if err != nil {
//		http.Error(w, "invalid credentials", http.StatusUnauthorized)
//		return
//	}
//
//	//we need to compare the password between what the user provided and the one saved on the database
//	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
//		http.Error(w, "invalid credentials", http.StatusUnauthorized)
//		return
//	}
//
//	tokenService := service.NewTokenService(h.accessConfig, h.refreshConfig)
//
//	accessString, err := tokenService.GenerateAccessToken(user.ID)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	refreshString, err := tokenService.GenerateRefreshToken(user.ID)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	resp := endpoint.LoginResponse{
//		AccessToken:  accessString,
//		RefreshToken: refreshString,
//	}
//
//	w.WriteHeader(http.StatusOK)
//	//encoding resp
//	if err := json.NewEncoder(w).Encode(resp); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	tokenService := service.NewTokenService(h.accessConfig, h.refreshConfig)

	//in postman headers, authorization, take its value and check if Bearer is there, and the 2nd part is the token which
	//we need to pass to our ValidateAccessToken()
	claims, err := tokenService.ValidateRefreshToken(service.GetTokenFromBearerString(
		r.Header.Get("Authorization")))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	//The claims.ID is the user ID
	accessString, err := tokenService.GenerateAccessToken(claims.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	refreshString, err := tokenService.GenerateRefreshToken(claims.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := endpoint.LoginResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *AuthHandler) CheckAccessTokenValidity(w http.ResponseWriter, r *http.Request) {
	tokenService := service.NewTokenService(h.accessConfig, h.refreshConfig)

	_, err := tokenService.ValidateAccessToken(service.GetTokenFromBearerString(
		r.Header.Get("Authorization")))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

}
