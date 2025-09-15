package repository

import (
	"eBlog/internal/db"
	"eBlog/internal/models"
)

func GetUserByUsername(username string) (models.User, error) {
	var u models.User
	const query = `SELECT id, username, password
					FROM users 
					WHERE username=$1`

	err := db.GetDBConnection().Get(&u, query, username)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func GetUserByUsernameAndPassword(username, password string) (models.User, error) {
	var u models.User
	const query = `SELECT id, username, password
					FROM users 
					WHERE username=$1 AND password=$2`

	err := db.GetDBConnection().Get(&u, query, username, password)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func CreateUser(u models.User) error {
	_, err := db.GetDBConnection().
		Exec("INSERT INTO users (username, password, full_name, address) VALUES ($1, $2, $3, $4)",
			u.Username, u.Password, u.FullName, u.Address)
	if err != nil {
		return err
	}

	return nil
}
