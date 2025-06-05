package article

import (
	"context"
	"golang/golang_clean_architecture/internal/models"
	"golang/golang_clean_architecture/internal/repository"
)

type UseCase struct {
	articleRepository repository.ArticleRepository
}

func New(repo repository.ArticleRepository) *UseCase {
	return &UseCase{articleRepository: repo}
}

func (u *UseCase) GetAll(ctx context.Context) ([]models.ArticleResponse, error) {
	var response []models.ArticleResponse

	res, err := u.articleRepository.GetAll(ctx)
	if err != nil {
		return response, err
	}

	for _, v := range res {
		response = append(response, v.ToArticleResponse())
	}

	return response, err
}
