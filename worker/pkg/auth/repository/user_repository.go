package repository

import (
	"database/sql"
	"errors"
	"log"

	userModel "github.com/jsusmachaca/tiksup/pkg/auth/model"
	"github.com/jsusmachaca/tiksup/pkg/auth/validation"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (user *UserRepository) InsertUser(data userModel.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO 
	users(first_name, username, password) 
	VALUES ($1, $2, $3);`
	stmt, err := user.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(data.FirstName, data.Username, string(bytes))
	if err != nil {
		log.Println(err)
		return err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if i != 1 {
		return errors.New("1 row was expected to be affect")
	}
	log.Println("user inserted success")
	return nil
}

func (user *UserRepository) GetUser(data userModel.User) (userModel.User, error) {
	var userData userModel.User

	query := `SELECT id, username, password
	FROM users WHERE username=$1;`
	rows, err := user.DB.Query(query, data.Username)
	if err != nil {
		log.Println(err)
		return userData, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(
			&userData.ID,
			&userData.Username,
			&userData.Password,
		); err != nil {
			return userData, err
		}
	} else {
		return userData, validation.ErrIncorrectCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password))
	if err != nil {
		return userData, validation.ErrIncorrectCredentials
	}

	return userData, nil
}
