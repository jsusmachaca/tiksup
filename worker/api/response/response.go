package response

import "github.com/jsusmachaca/tiksup/pkg/movie/model"

type RandoMovie struct {
	UserID string        `json:"user_id"`
	Movies []model.Movie `json:"movies"`
}
