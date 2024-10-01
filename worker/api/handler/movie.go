package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jsusmachaca/tiksup/internal/util"
	"github.com/jsusmachaca/tiksup/pkg/movie/repository"
)

func GetMovieScores(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	movie := repository.MovieRository{DB: db}

	token := r.Header.Get("Authorization")

	if !strings.HasPrefix(token, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Token not provided"}`))
		return
	}
	token = token[7:]
	claims, err := util.ValidateToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Token is not valid"}`))
		return
	}

	recomendation, err := movie.GetPreferences(claims["user_id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Internal server error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recomendation); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Internal server error"}`))
		return
	}
}
