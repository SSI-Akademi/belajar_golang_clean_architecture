package repository

import (
	"context"
	"golang/golang_clean_architecture/internal/models"
)

type ArticleRepository interface {
	GetAll(ctx context.Context) ([]*models.Article, error)
}
