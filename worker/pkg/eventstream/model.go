package eventstream

type KafkaData struct {
	UserID      string      `json:"user_id"`
	MovieID     string      `json:"movie_id"`
	Preferences Preferences `json:"preferences"`
	Next        bool        `json:"next"`
}

type Preferences struct {
	GenreScore       []GenreScore     `json:"genre_score"`
	ProtagonistScore ProtagonistScore `json:"protagonist_score"`
	DirectorScore    DirectorScore    `json:"director_score"`
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
