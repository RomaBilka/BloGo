package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	User     string
	Password string
	Database string
	Host     string
}

func Run(config Config) *sql.DB {
	db := connect(config)
	return db
}

func connect(config Config) *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", config.User, config.Password, config.Host, config.Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}
