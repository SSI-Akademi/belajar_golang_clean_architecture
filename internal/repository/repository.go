package repository

import (
	"context"
	"golang/golang_clean_architecture/internal/models"
)

type ArticleRepository interface {
	GetAll(ctx context.Context) ([]*models.Article, error)
	Store(ctx context.Context, article *models.Article) (int64, error)
	Delete(ctx context.Context, id int64) (bool, error)
	Update(ctx context.Context, article *models.Article) (*models.Article, error)
}
