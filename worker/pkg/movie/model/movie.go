package model

type Movie struct {
	URL         string   `json:"url"`
	Title       string   `json:"title"`
	Genre       []string `json:"genre"`
	Protagonist string   `json:"protagonist"`
	Director    string   `json:"director"`
}

type MovieRemendation struct {
	UserID      string `json:"user_id"`
	Preferences `json:"preferences"`
	Movies      []Movie `json:"movies"`
}

type Preferences struct {
	GenreScore       []GenreScore       `json:"genre_score"`
	ProtagonistScore []ProtagonistScore `json:"protagonist_score"`
	DirectorScore    []DirectorScore    `json:"director_score"`
}

type GenreScore struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type ProtagonistScore struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type DirectorScore struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type History struct {
	UserID  string `json:"user_id"`
	MovieID string `json:"movie_id"`
}
