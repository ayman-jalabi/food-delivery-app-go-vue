package handlers

import (
	"encoding/json"
	"io"
	"main/endpoint"
	"main/repos"
	"net/http"
)

type UserHandler struct {
	Repo *repos.UserRepo
}

func NewUserHandler(repo *repos.UserRepo) UserHandler {
	return UserHandler{
		Repo: repo,
	}
}

func (uh UserHandler) CheckIfEmailExists(writer http.ResponseWriter, request *http.Request) {
	// Read the request body
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	// Unmarshal the JSON into the struct
	var user endpoint.RegistrationRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	exists, err := uh.Repo.CheckEmailExistence(user.Email)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	if exists {
		// Email already exists
		http.Error(writer, "Email already registered, please use another email", http.StatusConflict)
		return
	}
}

//func (uh UserHandler) GetUserInfo(writer http.ResponseWriter, _ *http.Request) {
//	users, err := uh.Repo.GetUser(email)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	byteSlice, err := json.Marshal(users)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	writer.Header().Set("Content-Type", "application/json")
//	writer.Write(byteSlice)
//}
//
//func (uh UserHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {
//	userToCreate := models.UserJson{}
//
//	err := json.NewDecoder(request.Body).Decode(&userToCreate)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userToCreate.Password), bcrypt.MinCost)
//	userToCreate.Password = string(hashedPassword)
//
//	userModel := models.User{
//		ID:          "",
//		Email:       userToCreate.Email,
//		FirstName:   userToCreate.FirstName,
//		LastName:    userToCreate.LastName,
//		Address:     userToCreate.Address,
//		PhoneNumber: userToCreate.PhoneNumber,
//		Password:    userToCreate.Password,
//	}
//
//	err = uh.Repo.CreateUser(userModel)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	writer.WriteHeader(http.StatusCreated)
//}
