package eventstream

import (
	"database/sql"
	"log"

	"github.com/jsusmachaca/tiksup/pkg/auth"
)

type KafkaRepository struct {
	DB *sql.DB
}

func (kafka *KafkaRepository) UpdateUserInfo(data KafkaData) error {
	auth := auth.UserRepository{DB: kafka.DB}

	preferenceID, err := auth.GetPreferenceID(data.UserID)
	if err != nil {
		return err
	}

	tx, err := kafka.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			log.Println("Kafka transaction rolled back:", err)
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	queryGenre := `INSERT INTO genre_score 
		(preference_id, name, score) 
		VALUES ($1, $2, $3)
		ON CONFLICT (preference_id, name)
		DO UPDATE SET score = genre_score.score + EXCLUDED.score;`
	stmtGenre, err := tx.Prepare(queryGenre)
	if err != nil {
		return err
	}
	defer stmtGenre.Close()

	for _, d := range data.Preferences.GenreScore {
		_, err := stmtGenre.Exec(preferenceID, d.Name, d.Score)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("Updated genre success")
	}

	queryProtagonist := `INSERT INTO protagonist_score 
			(preference_id, name, score) 
			VALUES ($1, $2, $3)
			ON CONFLICT (preference_id, name)
			DO UPDATE SET score = protagonist_score.score + EXCLUDED.score;`
	stmtProtagonist, err := tx.Prepare(queryProtagonist)
	if err != nil {
		return err
	}
	defer stmtProtagonist.Close()

	_, err = stmtProtagonist.Exec(
		preferenceID,
		data.Preferences.ProtagonistScore.Name,
		data.Preferences.ProtagonistScore.Score,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Updated protagonist success")

	queryDirector := `INSERT INTO director_score 
			(preference_id, name, score) 
			VALUES ($1, $2, $3)
			ON CONFLICT (preference_id, name)
			DO UPDATE SET score = director_score.score + EXCLUDED.score;`
	stmtDirector, err := tx.Prepare(queryDirector)
	if err != nil {
		return err
	}
	defer stmtDirector.Close()

	_, err = stmtDirector.Exec(
		preferenceID,
		data.Preferences.DirectorScore.Name,
		data.Preferences.DirectorScore.Score,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Updated director success")

	history := `INSERT INTO history (user_id, movie_id) 
		VALUES ($1, $2)
		ON CONFLICT (user_id, movie_id)
		DO NOTHING;`
	stmtHistory, err := tx.Prepare(history)
	if err != nil {
		return err
	}
	defer stmtHistory.Close()

	movie_id := IsValidObjectID(data.MovieID)
	_, err = stmtHistory.Exec(data.UserID, movie_id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Insert movie in history")

	return nil
}
