package controllers

import (
	"api/src/database"
	"api/src/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser handles the user create route
func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error = json.Unmarshal(reqBody, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}
}

// ListUsers retrieves all users from database
func ListUsers(w http.ResponseWriter, r *http.Request) {

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
