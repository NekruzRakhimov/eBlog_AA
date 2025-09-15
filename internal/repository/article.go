package repository

import (
	"database/sql"
	"eBlog/internal/db"
	"eBlog/internal/models"
	"fmt"
)

func CreateArticle(a models.Article) error {
	_, err := db.GetDBConnection().
		Exec("INSERT INTO articles (title, description, user_id) VALUES ($1, $2, $3)", a.Title, a.Description, a.UserID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllArticles(userID int, title string) ([]models.Article, error) {
	var articles []models.Article
	title = "'%" + title + "%'" // %sdfkdlf%
	var query = fmt.Sprintf(`
		SELECT
			id, 
			title, 
			description,
			user_id,
			created_at
		FROM articles 
		WHERE deleted_at IS NULL 
		  AND user_id = $1 
			AND title ILIKE %s
		ORDER BY created_at DESC`, title)

	err := db.GetDBConnection().Select(&articles, query, userID)
	if err != nil {
		return []models.Article{}, err
	}
	return articles, nil
}

func GetArticleByID(id int) (models.Article, error) {
	var article models.Article
	const query = `SELECT id, 
       title, 
       description, 
       user_id,
       created_at
					FROM articles 
					WHERE deleted_at IS NULL AND id=$1`

	err := db.GetDBConnection().Get(&article, query, id)
	if err != nil {
		return models.Article{}, err
	}
	return article, nil
}

func UpdateArticle(a models.Article) error {
	query := `UPDATE articles SET
					title = $1,
					description = $2,
					updated_at = CURRENT_TIMESTAMP
				WHERE id = $3`

	result, err := db.GetDBConnection().Exec(query, a.Title, a.Description, a.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func DeleteArticle(id int) error {
	query := "UPDATE articles SET deleted_at = NOW() WHERE id = $1"

	result, err := db.GetDBConnection().Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
