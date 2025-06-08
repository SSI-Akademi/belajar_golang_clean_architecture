package article

import (
	"context"
	"errors"
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

	return response, nil
}

func (u *UseCase) Store(ctx context.Context, request models.ArticleCreateRequest) (models.ArticleResponse, error) {
	var article = new(models.Article)
	var response models.ArticleResponse

	// convert request to article struct
	article.FromCreateRequest(request)

	// executed store
	articleID, err := u.articleRepository.Store(ctx, article)
	if err != nil {
		return response, err
	}

	article.ID = articleID
	response = article.ToArticleResponse()
	return response, nil
}

func (u *UseCase) Delete(ctx context.Context, id int64) (bool, error) {
	var isSuccess bool

	isSuccess, err := u.articleRepository.Delete(ctx, id)
	if err != nil {
		return isSuccess, err
	}

	return isSuccess, nil
}

func (u *UseCase) Update(ctx context.Context, request models.ArticleUpdateRequest) (models.ArticleResponse, error) {
	var article = new(models.Article)
	var response models.ArticleResponse

	if request.ID == 0 {
		return response, errors.New("request article_id do not zero or empty")
	}

	// from request to article struct
	article.FromUpdateRequest(request)

	// execute update
	article, err := u.articleRepository.Update(ctx, article)
	if err != nil {
		return response, err
	}

	// append data update to response
	response = article.ToArticleResponse()
	return response, nil
}
