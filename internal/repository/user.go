package repository

import (
	"eBlog/internal/models"
)

func (r *Repository) GetUserByUsername(username string) (models.User, error) {
	var u models.User
	const query = `SELECT id, username, password
					FROM users 
					WHERE username=$1`

	err := r.db.Get(&u, query, username)
	if err != nil {
		return models.User{}, r.translateError(err)
	}
	return u, nil
}

func (r *Repository) GetUserByUsernameAndPassword(username, password string) (models.User, error) {
	var u models.User
	const query = `SELECT id, username, password
					FROM users 
					WHERE username=$1 AND password=$2`

	err := r.db.Get(&u, query, username, password)
	if err != nil {
		return models.User{}, r.translateError(err)
	}
	return u, nil
}

func (r *Repository) CreateUser(u models.User) error {
	_, err := r.db.
		Exec("INSERT INTO users (username, password, full_name, address) VALUES ($1, $2, $3, $4)",
			u.Username, u.Password, u.FullName, u.Address)
	if err != nil {
		return r.translateError(err)
	}

	return nil
}
