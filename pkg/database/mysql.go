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

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		fmt.Println("test")
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("test 2")
		log.Fatal(err)
	}

	fmt.Println("Database connect succesfully")
	return db, err
}
