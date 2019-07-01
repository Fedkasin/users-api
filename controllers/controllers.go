package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/Fedkasin/users-api/models"
	userService "github.com/Fedkasin/users-api/services"
)

func GetUserByIdController(res http.ResponseWriter, req *http.Request) {
	var user models.User

	res.Header().Add("Content-Type", "application/json")

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	
	err := userService.GetUserById(id, &user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(user)
}

func UpdateUserByIdController(res http.ResponseWriter, req *http.Request) {
	var user models.User

	res.Header().Add("Content-Type", "application/json")

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	_ = json.NewDecoder(req.Body).Decode(&user)

	result := userService.UpdateUserById(id, user)
	json.NewEncoder(res).Encode(result)
}

func RemoveUserByIdController(res http.ResponseWriter, req *http.Request) {

}

func CreateUserController(res http.ResponseWriter, req *http.Request) {
	var user models.User

	res.Header().Add("Content-Type", "application/json")
	_ = json.NewDecoder(req.Body).Decode(&user)

	createdUser, err := userService.CreateUser(user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(createdUser)
}

func GetAllUsersController(res http.ResponseWriter, req *http.Request) {
	var allUsers []models.User

	res.Header().Add("Content-Type", "application/json")

	result, err := userService.GetUsers(allUsers)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(result)
}
