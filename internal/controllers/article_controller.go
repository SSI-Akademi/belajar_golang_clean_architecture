package controllers

import (
	"golang/golang_clean_architecture/internal/models"
	"golang/golang_clean_architecture/internal/usecase"
	"golang/golang_clean_architecture/pkg/util"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	articleUsecase usecase.ArticleUsecase
	validate       *validator.validate
}

func New(articleUsecase usecase.ArticleUsecase) *Controller {
	return &Controller{
		articleUsecase: articleUsecase,
		validate:       validator.new(),
	}
}

func (d *Controller) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleResponse
	var statusCode int = http.StatusBadRequest

	defer func () {
		util.Response(w, response, statusCode)
	}

	res, err := d.articleUsecase.GetAll(r.Context())
	if err != nil {
		log.Fatal(err)
	}

	statusCode = http.StatusOK
	response.Data = append(response.Data, res...)
	response.Code = statusCode
	response.Status = "OK"
}
