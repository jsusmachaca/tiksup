package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/jsusmachaca/tiksup/api/response"
	"github.com/jsusmachaca/tiksup/internal/util"
	modelUser "github.com/jsusmachaca/tiksup/pkg/auth/model"
	userRepository "github.com/jsusmachaca/tiksup/pkg/auth/repository"
	"github.com/jsusmachaca/tiksup/pkg/auth/validation"
)

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer r.Body.Close()
	user := userRepository.UserRepository{DB: db}

	w.Header().Set("Content-Type", "application/json")

	var body modelUser.User
	if err := validation.UserValidation(r.Body, &body); err != nil {
		response := response.ErrorResponse{Error: "Invalid data"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := user.GetUser(body)
	if err != nil {
		response := response.ErrorResponse{Error: "Invalid username or password"}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	token, err := util.CreateToken(data.ID, data.Username)
	if err != nil {
		response := response.ErrorResponse{Error: "Internal server error"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]string{
		"access_token": token,
	}); err != nil {
		response := response.ErrorResponse{Error: "Internal server error"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer r.Body.Close()
	user := userRepository.UserRepository{DB: db}

	w.Header().Set("Content-Type", "application/json")

	var body modelUser.User
	if err := validation.UserValidation(r.Body, &body); err != nil {
		response := response.ErrorResponse{Error: "Invalid data"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err := user.InsertUser(body)
	if err != nil {
		response := response.ErrorResponse{Error: "Error registering user"}
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = user.CreatePreference(body)
	if err != nil {
		response := response.ErrorResponse{Error: "Internal server error"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	successResponse := map[string]string{
		"first_name": body.FirstName,
		"username":   body.Username,
		"password":   body.Password,
	}
	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		response := response.ErrorResponse{Error: "Internal server error"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
}
