package usecase

import (
	"context"
	"golang/golang_clean_architecture/internal/models"
)

type ArticleUsecase interface {
	GetAll(ctx context.Context) ([]models.ArticleResponse, error)
	Store(ctx context.Context, article models.ArticleCreateRequest) (models.ArticleResponse, error)
	Delete(ctx context.Context, id int64) (bool, error)
	Update(ctx context.Context, article models.ArticleUpdateRequest) (models.ArticleResponse, error)
}
