package main

import (
	"golang/golang_clean_architecture/pkg/wire"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server, err := wire.InitializedServer()
	if err != nil {
		log.Fatal(err)
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
