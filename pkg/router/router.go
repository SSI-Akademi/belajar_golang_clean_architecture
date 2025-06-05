package router

import (
	"encoding/json"
	"golang/golang_clean_architecture/internal/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewServer(router *httprouter.Router) *http.Server {
	// implementasi server di sini
	return &http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}
}

func NewRouter(handler *controllers.Controller) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		json.NewEncoder(w).Encode("Test")
	})
	router.GET("/articles", handler.GetAll)

	return router
}
