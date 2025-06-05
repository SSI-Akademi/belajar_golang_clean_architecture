package usecase

import (
	"context"
	"golang/golang_clean_architecture/internal/models"
)

type ArticleUsecase interface {
	GetAll(ctx context.Context) ([]*models.ArticleResponse, error)
}
