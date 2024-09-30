package repository

import (
	"database/sql"
	"log"

	"github.com/jsusmachaca/tiksup/pkg/auth/repository"
	"github.com/jsusmachaca/tiksup/pkg/movie/model"
)

type MovieRository struct {
	DB *sql.DB
}

func (movie *MovieRository) GetPreferences(user_id string) (model.MovieRemendation, error) {
	var recomendation model.MovieRemendation
	var genre model.GenreScore
	var protagonist model.ProtagonistScore
	var director model.DirectorScore

	user := repository.UserRepository{DB: movie.DB}

	recomendation.UserID = user_id

	preferenceID, err := user.GetPreferenceID(user_id)
	if err != nil {
		return recomendation, err
	}

	tx, err := movie.DB.Begin()
	if err != nil {
		return recomendation, err
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
		return recomendation, err
	}
	defer rowsGenre.Close()

	for rowsGenre.Next() {
		if err := rowsGenre.Scan(
			&genre.Name,
			&genre.Score,
		); err != nil {
			return recomendation, err
		}
		recomendation.Preferences.GenreScore = append(recomendation.Preferences.GenreScore, genre)
	}

	// Protagonist Query
	queryProtagonist := `SELECT name, score FROM preference
		JOIN protagonist_score 
		ON protagonist_score.preference_id=preference.id
		WHERE preference_id=$1;`
	rowsProtagonist, err := tx.Query(queryProtagonist, preferenceID)
	if err != nil {
		return recomendation, err
	}
	defer rowsProtagonist.Close()

	for rowsProtagonist.Next() {
		if err := rowsProtagonist.Scan(
			&protagonist.Name,
			&protagonist.Score,
		); err != nil {
			return recomendation, err
		}
		recomendation.Preferences.ProtagonistScore = append(recomendation.Preferences.ProtagonistScore, protagonist)
	}

	// Director Query
	queryDirector := `SELECT name, score FROM preference
		JOIN director_score 
		ON director_score.preference_id=preference.id
		WHERE preference_id=$1;`
	rowsDirector, err := tx.Query(queryDirector, preferenceID)
	if err != nil {
		return recomendation, err
	}
	defer rowsDirector.Close()

	for rowsDirector.Next() {
		if err := rowsDirector.Scan(
			&director.Name,
			&director.Score,
		); err != nil {
			return recomendation, err
		}
		recomendation.Preferences.DirectorScore = append(recomendation.Preferences.DirectorScore, director)
	}

	return recomendation, nil
}

func (movie *MovieRository) GetHistory(user_id string) ([]model.History, error) {
	var history model.History
	var historyArray []model.History

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
		historyArray = append(historyArray, history)
	}

	return historyArray, nil
}
