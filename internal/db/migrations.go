package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	createUsersTableDDL = `CREATE TABLE IF NOT EXISTS users (
    	id          SERIAL PRIMARY KEY,
    	username    VARCHAR(255) NOT NULL UNIQUE,
    	password    VARCHAR(255) NOT NULL,
    	full_name   VARCHAR(255) NOT NULL,
    	address     VARCHAR(255) NOT NULL,
    	created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at  TIMESTAMP
	);`

	createArticlesTableDDL = `
	CREATE TABLE IF NOT EXISTS articles
	(
		id          SERIAL PRIMARY KEY,
		title       VARCHAR(255) NOT NULL UNIQUE,
		description TEXT NOT NULL,
		user_id     INT REFERENCES users(id),
		created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at  TIMESTAMP,
		deleted_at  TIMESTAMP DEFAULT NULL
	);`
)

func RunMigrations(db *sqlx.DB) error {
	_, err := db.Exec(createUsersTableDDL)
	if err != nil {
		fmt.Println("error creating users table")
		return err
	}

	_, err = db.Exec(createArticlesTableDDL)
	if err != nil {
		fmt.Println("error creating articles table")
		return err
	}

	return nil
}
