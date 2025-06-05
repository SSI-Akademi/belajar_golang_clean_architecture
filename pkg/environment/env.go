package environment

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() (string, error) {
	fileEnv := ".env"
	err := godotenv.Load(fileEnv)

	if err != nil {
		log.Fatal(err)
	}

	return fileEnv, err
}
