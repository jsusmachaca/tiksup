package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jsusmachaca/tiksup/api/response"
	"github.com/jsusmachaca/tiksup/internal/util"
	"github.com/jsusmachaca/tiksup/pkg/movie"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	movie := movie.MovieRepository{DB: db}

	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")
	if !strings.HasPrefix(token, "Bearer ") {
		response.WriteJsonError(w, "Token not provided", http.StatusUnauthorized)
		return
	}
	token = token[7:]
	claims, err := util.ValidateToken(token)
	if err != nil {
		response.WriteJsonError(w, "Token is not valid", http.StatusUnauthorized)
		return
	}

	recomendation, err := movie.GetPreferences(claims["user_id"].(string))
	if err != nil {
		response.WriteJsonError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(recomendation); err != nil {
		response.WriteJsonError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func GetRandomMovies(w http.ResponseWriter, r *http.Request, db *sql.DB, mongoConn movie.MongoConnection) {
	movieMongo := mongoConn.ToRepository()
	var randomMovie []movie.Movie

	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")
	if !strings.HasPrefix(token, "Bearer ") {
		response.WriteJsonError(w, "Token not provided", http.StatusUnauthorized)
		return
	}
	token = token[7:]
	claims, err := util.ValidateToken(token)
	if err != nil {
		response.WriteJsonError(w, "Token is not valid", http.StatusUnauthorized)
		return
	}

	err = movieMongo.GetRadomMovies(&randomMovie)
	if err != nil {
		response.WriteJsonError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	movieResponse := response.RandoMovie{
		UserID: claims["user_id"].(string),
		Movies: randomMovie,
	}

	if err := json.NewEncoder(w).Encode(movieResponse); err != nil {
		response.WriteJsonError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
