package handlers

import (
	"encoding/json"
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var auth models.Login

	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		
	}
	json.NewEncoder(w).Encode(checkLogin(auth))
}

func checkLogin(auth models.Login) string {
	if auth.Username != models.LogUser.Password || auth.Password != models.LogUser.Password {
		return "Invalid credentials"
	}

	validToken, err := utils.GenerateJWT()
	if err != nil {
		return "Error generate token."
	}
	return validToken
}
