package router

import (
	"golang/golang_clean_architecture/internal/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}
}

func NewRouter(handler *controllers.Controller) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/articles", handler.GetAll)
	router.POST("/api/articles", handler.Store)
	router.DELETE("/api/articles/:article_id", handler.Delete)
	router.PUT("/api/articles/:article_id", handler.Update)

	return router
}
