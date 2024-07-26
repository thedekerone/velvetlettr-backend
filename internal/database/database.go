package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func InitDB(config Config) error {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s port=%d password=%s", config.User, config.DBName, config.SSLMode, config.Port, config.Password)
	log.Print(connStr)
	db, connError := sql.Open("postgres", connStr)

	if connError != nil {
		log.Fatal(connError)
		return connError
	}

	err := db.Ping()

	if err != nil {
		const logError = "Error when trying to open connection to db"
		log.Fatal(err)
		return fmt.Errorf(logError)
	}
	DB = db
	log.Print("Database initialized correctly")
	return nil
}

func CloseDB() error {
	if DB == nil {
		return fmt.Errorf("DB has not been initialized")
	}

	err := DB.Close()

	if err != nil {
		return err
	}

	return nil
}
