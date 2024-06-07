package handlers

import (
	"encoding/json"
	"migration/config"
	"migration/models"
	"net/http"
)

// CreateUserHandler handles the creation of a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// GetUsersHandler handles fetching all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	config.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}
