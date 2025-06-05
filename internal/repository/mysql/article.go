package mysql

import (
	"context"
	"database/sql"
	"golang/golang_clean_architecture/internal/models"
)

type ArticleStore struct {
	db *sql.DB
}

func New(conn *sql.DB) *ArticleStore {
	return &ArticleStore{conn}
}

func (r *ArticleStore) GetAll(ctx context.Context) ([]*models.Article, error) {
	var result []*models.Article

	const querySelect = "SELECT id, title, content, create_at, updated_at FROM articles"
	rows, err := r.db.QueryContext(ctx, querySelect)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		t := models.Article{}
		err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.CreateAt, &t.UpdateAt)

		if err != nil {
			return result, err
		}

		result = append(result, &t)
	}

	return result, nil
}
