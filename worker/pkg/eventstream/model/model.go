package model

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

/*
Model
{
    "user_id": 1213,
    "video_id": 12312,
    "waching_time": 12s,
    "waching_repeat": 2,
    "preferences": {
        "genre_scores": [
            {
		"name": "terror",
		"score": 1
	    },
	    {
		"name": "fantacy",
		"score": 1
	    },
	    {
		"name": "action",
		"score": 1
	    },
            ...
        ],
        "actors_scores": {
		"name": "Leo DiCaprio",
		"score": 1
	},
        "director_scores": {
		"name": "Cuentin Tarantino",
		"score": 4
	}
    },
    "next": false
}
*/
