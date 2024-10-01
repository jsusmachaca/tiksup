package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URL         string             `bson:"url" json:"url"`
	Title       string             `bson:"title" json:"title"`
	Genre       []string           `bson:"genre" json:"genre"`
	Protagonist string             `bson:"protagonist" json:"protagonist"`
	Director    string             `bson:"director" json:"director"`
}

type MovieRemmendation struct {
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
