package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jsusmachaca/go-router/pkg/response"
	"github.com/jsusmachaca/tiksup/api/middleware"
	"github.com/jsusmachaca/tiksup/api/model"
	"github.com/jsusmachaca/tiksup/pkg/movie"
)

type GetUserInfo struct {
	DB *sql.DB
}

type GetRandomMovies struct {
	DB        *sql.DB
	MongoConn movie.MongoConnection
}

func (h *GetUserInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	movie := &movie.MovieRepository{DB: h.DB}

	w.Header().Set("Content-Type", "application/json")
	claims, ok := r.Context().Value(middleware.TokenClaims).(jwt.MapClaims)
	if !ok {
		response.JsonErrorFromString(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	recomendation, err := movie.GetPreferences(claims["user_id"].(string))
	if err != nil {
		response.JsonErrorFromString(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(recomendation); err != nil {
		response.JsonErrorFromString(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func (h *GetRandomMovies) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	movieMongo := h.MongoConn.ToRepository()
	var randomMovie []movie.Movie

	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.TokenClaims).(jwt.MapClaims)
	if !ok {
		response.JsonErrorFromString(w, "Internal server error", http.StatusInternalServerError)
	}

	err := movieMongo.GetRadomMovies(&randomMovie)
	if err != nil {
		response.JsonErrorFromString(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	movieResponse := model.RandoMovie{
		UserID: claims["user_id"].(string),
		Movies: randomMovie,
	}

	response.JsonResponse(w, movieResponse, 200)
}
