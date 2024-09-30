package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/jsusmachaca/tiksup/api/response"
	"github.com/jsusmachaca/tiksup/internal/util"
	modelUser "github.com/jsusmachaca/tiksup/pkg/auth/model"
	"github.com/jsusmachaca/tiksup/pkg/auth/repository"
	"github.com/jsusmachaca/tiksup/pkg/auth/validation"
)

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer r.Body.Close()

	user := repository.UserRepository{DB: db}

	var body modelUser.User

	if err := validation.UserValidation(r.Body, &body); err != nil {
		response := response.ErrorResponse{Error: "Invalid data"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := user.GetUser(body)
	if err != nil {
		response := response.ErrorResponse{Error: "Invalid username or password"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	token, err := util.CreateToken(data.ID, data.Username)
	if err != nil {
		response := response.ErrorResponse{Error: "Internal server error"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": token,
	})
}

func Register(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer r.Body.Close()

	user := repository.UserRepository{DB: db}

	var body modelUser.User

	if err := validation.UserValidation(r.Body, &body); err != nil {
		response := response.ErrorResponse{Error: "Invalid data"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err := user.InsertUser(body)
	if err != nil {
		response := response.ErrorResponse{Error: "Error registering user"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = user.CreatePreference(body)
	if err != nil {
		response := response.ErrorResponse{Error: "Internal server error"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"first_name": body.FirstName,
		"username":   body.Username,
		"password":   body.Password,
	}

	json.NewEncoder(w).Encode(response)
}
