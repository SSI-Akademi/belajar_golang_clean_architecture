package controllers

import (
	"encoding/json"
	"errors"
	"golang/golang_clean_architecture/internal/models"
	"golang/golang_clean_architecture/internal/usecase"
	"golang/golang_clean_architecture/pkg/util"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	articleUsecase usecase.ArticleUsecase
	validate       *validator.Validate
}

func New(articleUsecase usecase.ArticleUsecase) *Controller {
	return &Controller{
		articleUsecase: articleUsecase,
		validate:       validator.New(),
	}
}

func (d *Controller) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var statusCode int = http.StatusBadRequest

	defer func() {
		util.Response(w, response, statusCode)
	}()

	res, err := d.articleUsecase.GetAll(r.Context())
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	statusCode = http.StatusOK
	response.Data = append(response.Data, res...)
	response.Code = statusCode
	response.Status = "OK"

}

func (d *Controller) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var request models.ArticleCreateRequest
	var statusCode int = http.StatusBadRequest

	defer func() {
		util.Response(w, response, statusCode)
	}()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	err = d.validate.Struct(request)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	res, err := d.articleUsecase.Store(r.Context(), request)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	response.Data = append(response.Data, res)
	statusCode = http.StatusOK
	response.Code = statusCode
	response.Status = "OK"
}

func (d *Controller) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var statusCode int = http.StatusBadRequest

	defer func() {
		util.Response(w, response, statusCode)
	}()

	articleIDString := params.ByName("article_id")
	articleID, err := strconv.ParseInt(articleIDString, 0, 64)
	res, err := d.articleUsecase.Delete(r.Context(), articleID)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	if !res {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = errors.New("unknown error").Error()
		return
	}

	statusCode = http.StatusOK
	response.Code = statusCode
	response.Status = "OK"
}

func (d *Controller) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response models.ArticleListResponse
	var request models.ArticleUpdateRequest
	var statusCode int = http.StatusBadRequest

	defer func() {
		util.Response(w, response, statusCode)
	}()

	articleIDString := params.ByName("article_id")
	articleID, err := strconv.ParseInt(articleIDString, 0, 64)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	request.ID = articleID

	err = d.validate.Struct(request)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	res, err := d.articleUsecase.Update(r.Context(), request)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.Code = statusCode
		response.Status = err.Error()
		return
	}

	statusCode = http.StatusOK
	response.Data = append(response.Data, res)
	response.Code = statusCode
	response.Status = "OK"
}
