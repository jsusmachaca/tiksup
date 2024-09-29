package model

type KafkaData struct {
	UserId         string      `json:"user_id"`
	VideoId        string      `json:"vide_id"`
	WatchingTime   string      `json:"waching_time"`
	WatchingRepeat string      `json:"waching_repeat"`
	Preferences    Preferences `json:"preferences"`
	Next           bool        `json:"next"`
}

type Preferences struct {
	GenreScores   []GenreScores    `json:"genre_scores"`
	Actorscores   []ActorScores    `json:"actor_scores"`
	Directorcores []DirectorScores `json:"director_scores"`
}

type GenreScores struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type ActorScores struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type DirectorScores struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
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
        "actors_scores": [
             {
		"name": "Leo DiCaprio",
		"score": 1
	     },
             {
		"name": "Brad Pitt",
		"score": 6
	     },
             {
		"name": "Angelina Jolie",
		"score": 4
	     },
		...
        ],
        "director_scores": [
             {
		"name": "Martin Scorsese",
		"score": 4
	     },
             {
		"name": "Cuentin Tarantino",
		"score": 4
	     },
             {
		"name": "Guillermo Del Toro",
		"score": 4
	     },
	     ...
        ]
    },
    "next": false
}
*/
