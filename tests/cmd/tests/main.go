package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/RomaBiliak/BloGo/project-layout/tests/internal/handlers"
	"github.com/RomaBiliak/BloGo/project-layout/tests/internal/repositories"
	"github.com/RomaBiliak/BloGo/project-layout/tests/internal/services"
	"github.com/RomaBiliak/BloGo/project-layout/tests/pkg/database/postgres"
	httpServer "github.com/RomaBiliak/BloGo/project-layout/tests/pkg/http"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	pgUser, ok := os.LookupEnv("PG_USER")
	if !ok {
		panic(errors.New("PG_USER is empty"))
	}
	pgPassword, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		panic(errors.New("PG_PASSWORD is empty"))
	}
	pgHost, ok := os.LookupEnv("PG_HOST")
	if !ok {
		panic(errors.New("Db Host is empty"))
	}
	pgDatabase, ok := os.LookupEnv("PG_DATABASE")
	if !ok {
		panic(errors.New("PG_DATABASE is empty"))
	}

	dbConfig := postgres.Config{
		pgUser,
		pgPassword,
		pgDatabase,
		pgHost,
	}
	db := postgres.Run(dbConfig)
	defer db.Close()

	userRepository := repositories.NewPostgreUserRepository(db)

	mux := http.NewServeMux()

	userService := services.NewUserService(userRepository)
	uHttp := handlers.NewUserHttp(userService)
	mux.HandleFunc("/create-user", uHttp.CreateUser)

	httpServer.Start(":8080", mux)
}
