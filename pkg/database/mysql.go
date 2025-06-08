package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func New(env string) (*sql.DB, error) {
	config := mysql.Config{
		User:                 os.Getenv("DATABASE_USER"),
		Passwd:               os.Getenv("DATABASE_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DATABASE_ADDRESS"),
		DBName:               os.Getenv("DATABASE_NAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open(os.Getenv("DATABASE_DRIVER"), config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database Connected")
	return db, err
}
