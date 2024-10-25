package response

import "github.com/jsusmachaca/tiksup/pkg/movie"

type RandoMovie struct {
	UserID string        `json:"user_id"`
	Movies []movie.Movie `json:"movies"`
}
