package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// CreateUser handles the user create route
func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewRepository(db)
	userID, err := repo.Create(user)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	user.ID = userID
	utils.JSON(w, http.StatusCreated, user)
}

// ListUsers retrieves all users from database
func ListUsers(w http.ResponseWriter, r *http.Request) {
	queryValue := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewRepository(db)
	users, err := repo.List(queryValue)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, users)
}

// GetUser returns one user by its id
func GetUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser changes information from a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// RemoveUser removes a user from the database
func RemoveUser(w http.ResponseWriter, r *http.Request) {

}
