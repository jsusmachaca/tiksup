package handler

import (
	"database/sql"
	"net/http"

	"github.com/jsusmachaca/go-router/pkg/response"
	"github.com/jsusmachaca/tiksup/internal/util"
	"github.com/jsusmachaca/tiksup/pkg/auth"
)

type Login struct {
	DB *sql.DB
}
type Register struct {
	DB *sql.DB
}

func (h *Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	user := &auth.UserRepository{DB: h.DB}

	w.Header().Set("Content-Type", "application/json")

	var body auth.User
	if err := auth.UserValidation(r.Body, &body); err != nil {
		response.JsonErrorFromString(w, "Invalid data", http.StatusBadRequest)
		return
	}

	data, err := user.GetUser(body)
	if err != nil {
		response.JsonErrorFromString(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := util.CreateToken(data.ID, data.Username)
	if err != nil {
		response.JsonErrorFromString(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response.JsonResponse(w, map[string]string{
		"access_token": token,
	}, 200)
}

func (h Register) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	user := &auth.UserRepository{DB: h.DB}

	w.Header().Set("Content-Type", "application/json")

	var body auth.User
	if err := auth.UserValidation(r.Body, &body); err != nil {
		response.JsonErrorFromString(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err := user.InsertUser(body)
	if err != nil {
		response.JsonErrorFromString(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	err = user.CreatePreference(body)
	if err != nil {
		response.JsonErrorFromString(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	successResponse := map[string]string{
		"first_name": body.FirstName,
		"username":   body.Username,
		"password":   body.Password,
	}
	response.JsonResponse(w, successResponse, 200)
}
