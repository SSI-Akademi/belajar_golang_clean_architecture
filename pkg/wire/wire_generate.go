package wire

import (
	"golang/golang_clean_architecture/pkg/database"
	"golang/golang_clean_architecture/pkg/environment"
	"golang/golang_clean_architecture/pkg/router"
	"net/http"
)

func InitializedServer() (*http.Server, error) {
	env, err := environment.Load()
	if err != nil {
		return nil, err
	}

	db, err := database.New(env)
	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, err
	}

	httpRouter := router.NewRouter()
	server := router.NewServer(httpRouter)

	return server, err
}
