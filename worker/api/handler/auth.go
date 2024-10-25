package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/jsusmachaca/tiksup/api/response"
	"github.com/jsusmachaca/tiksup/internal/util"
	"github.com/jsusmachaca/tiksup/pkg/auth"
)

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer r.Body.Close()
	user := auth.UserRepository{DB: db}

	w.Header().Set("Content-Type", "application/json")

	var body auth.User
	if err := auth.UserValidation(r.Body, &body); err != nil {
		response.WriteJsonError(w, "Invalid data", http.StatusBadRequest)
		return
	}

	data, err := user.GetUser(body)
	if err != nil {
		response.WriteJsonError(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := util.CreateToken(data.ID, data.Username)
	if err != nil {
		response.WriteJsonError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]string{
		"access_token": token,
	}); err != nil {
		response.WriteJsonError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer r.Body.Close()
	user := auth.UserRepository{DB: db}

	w.Header().Set("Content-Type", "application/json")

	var body auth.User
	if err := auth.UserValidation(r.Body, &body); err != nil {
		response.WriteJsonError(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err := user.InsertUser(body)
	if err != nil {
		response.WriteJsonError(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	err = user.CreatePreference(body)
	if err != nil {
		response.WriteJsonError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	successResponse := map[string]string{
		"first_name": body.FirstName,
		"username":   body.Username,
		"password":   body.Password,
	}
	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		response.WriteJsonError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
