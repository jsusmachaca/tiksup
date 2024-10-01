package repository

import (
	"database/sql"
	"log"

	"github.com/jsusmachaca/tiksup/pkg/auth/repository"
	"github.com/jsusmachaca/tiksup/pkg/movie/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieRository struct {
	DB *sql.DB
}

func (movie *MovieRository) GetPreferences(user_id string) (model.MovieRemmendation, error) {
	var recommendation model.MovieRemmendation
	recommendation.UserID = user_id
	recommendation.Preferences = model.Preferences{
		GenreScore:       []model.GenreScore{},
		ProtagonistScore: []model.ProtagonistScore{},
		DirectorScore:    []model.DirectorScore{},
	}
	user := repository.UserRepository{DB: movie.DB}

	preferenceID, err := user.GetPreferenceID(user_id)
	if err != nil {
		return recommendation, err
	}

	tx, err := movie.DB.Begin()
	if err != nil {
		return recommendation, err
	}

	defer func() {
		if err != nil {
			log.Println("Movies transaction rolled back:", err)
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Genre Query
	queryGenre := `SELECT name, score FROM preference
		JOIN genre_score 
		ON genre_score.preference_id=preference.id
		WHERE preference_id=$1;`
	rowsGenre, err := tx.Query(queryGenre, preferenceID)
	if err != nil {
		return recommendation, err
	}
	defer rowsGenre.Close()

	for rowsGenre.Next() {
		var genre model.GenreScore
		if err := rowsGenre.Scan(
			&genre.Name,
			&genre.Score,
		); err != nil {
			return recommendation, err
		}
		recommendation.Preferences.GenreScore = append(recommendation.Preferences.GenreScore, genre)
	}

	// Protagonist Query
	queryProtagonist := `SELECT name, score FROM preference
		JOIN protagonist_score 
		ON protagonist_score.preference_id=preference.id
		WHERE preference_id=$1;`
	rowsProtagonist, err := tx.Query(queryProtagonist, preferenceID)
	if err != nil {
		return recommendation, err
	}
	defer rowsProtagonist.Close()

	for rowsProtagonist.Next() {
		var protagonist model.ProtagonistScore
		if err := rowsProtagonist.Scan(
			&protagonist.Name,
			&protagonist.Score,
		); err != nil {
			return recommendation, err
		}
		recommendation.Preferences.ProtagonistScore = append(recommendation.Preferences.ProtagonistScore, protagonist)
	}

	// Director Query
	queryDirector := `SELECT name, score FROM preference
		JOIN director_score 
		ON director_score.preference_id=preference.id
		WHERE preference_id=$1;`
	rowsDirector, err := tx.Query(queryDirector, preferenceID)
	if err != nil {
		return recommendation, err
	}
	defer rowsDirector.Close()

	for rowsDirector.Next() {
		var director model.DirectorScore
		if err := rowsDirector.Scan(
			&director.Name,
			&director.Score,
		); err != nil {
			return recommendation, err
		}
		recommendation.Preferences.DirectorScore = append(recommendation.Preferences.DirectorScore, director)
	}

	return recommendation, nil
}

func (movie *MovieRository) GetHistory(user_id string) ([]primitive.ObjectID, error) {
	var history model.History
	var historyArray []primitive.ObjectID

	queryHistory := `SELECT movie_id FROM history WHERE user_id=$1;`
	rowsHistory, err := movie.DB.Query(queryHistory, user_id)
	if err != nil {
		return historyArray, err
	}
	defer rowsHistory.Close()

	for rowsHistory.Next() {
		if err := rowsHistory.Scan(
			&history.MovieID,
		); err != nil {
			return historyArray, err
		}
		objectID, err := primitive.ObjectIDFromHex(history.MovieID)
		if err != nil {
			return historyArray, err
		}
		historyArray = append(historyArray, objectID)
	}

	return historyArray, nil
}
