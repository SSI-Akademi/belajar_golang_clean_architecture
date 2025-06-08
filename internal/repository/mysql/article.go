package mysql

import (
	"context"
	"database/sql"
	"fmt"
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

	rows, err := r.db.QueryContext(ctx, "SELECT id, title, content, create_at, update_at FROM articles")
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

func (r *ArticleStore) Store(ctx context.Context, article *models.Article) (int64, error) {
	const queryInsert = "INSERT INTO articles(title, content, create_at, update_at) VALUES(?,?,?,?)"
	res, err := r.db.ExecContext(ctx, queryInsert, article.Title, article.Content, article.CreateAt, article.UpdateAt)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Printf("success create with lastId: %d", lastId)
	return lastId, nil
}

func (r *ArticleStore) Delete(ctx context.Context, id int64) (bool, error) {
	const queryDelete = "DELETE FROM articles WHERE id = ?"
	_, err := r.db.ExecContext(ctx, queryDelete, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ArticleStore) Update(ctx context.Context, article *models.Article) (*models.Article, error) {
	const queryUpdate = "UPDATE articles SET title = ?, content = ?, update_at = ? WHERE id = ?"
	res, err := r.db.ExecContext(ctx, queryUpdate, article.Title, article.Content, article.UpdateAt, article.ID)
	if err != nil {
		return nil, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	fmt.Println("Success updated at with affected %d", count)
	return article, nil
}
