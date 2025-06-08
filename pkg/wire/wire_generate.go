package wire

import (
	"golang/golang_clean_architecture/internal/controllers"
	"golang/golang_clean_architecture/internal/repository/mysql"
	"golang/golang_clean_architecture/internal/usecase/article"
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

	articleStore := mysql.New(db)
	articleUsecase := article.New(articleStore)
	delivery := controllers.New(articleUsecase)

	httpRouter := router.NewRouter(delivery)
	server := router.NewServer(httpRouter)

	return server, err
}
