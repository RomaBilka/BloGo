package repositories

import (
	"database/sql"
	"os"
	"testing"

	"github.com/RomaBiliak/BloGo/project-layout/tests/internal/models"
	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
)

var testUserRepository *UserRepository
var db *sql.DB

type test struct {
	Name string `faker:"name"`
	Email string `faker:"email"`
	Phone string `faker:"e_164_phone_number"`
}
var testData test

var userTest models.User

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	err = faker.FakeData(&userTest)
	if err != nil {
		panic(err)
	}

	userTest = models.User{Name: userTest.Name, Email: userTest.Email, Phone: userTest.Phone}
/*
	pgUser, ok := os.LookupEnv("PG_USER")
	if !ok {
		panic(errors.New("PG_USER is empty"))
	}
	pgPassword, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		panic(errors.New("PG_PASSWORD is empty"))
	}
	pgDatabase, ok := os.LookupEnv("PG_TEST_DATABASE")
	if !ok {
		panic(errors.New("PG_TEST_DATABASE is empty"))
	}
	pgHost, ok := os.LookupEnv("PG_HOST")
	if !ok {
		panic(errors.New("Db Host is empty"))
	}

	dbConfig := postgres.Config{
		pgUser,
		pgPassword,
		pgDatabase,
		pgHost,
	}
	db = postgres.Run(dbConfig)

	defer db.Close()
	testUserRepository = NewPostgreUserRepository(db)*/

	os.Exit(m.Run())
}

func truncateUsers() error {
	_, err := db.Query("TRUNCATE users CASCADE")
	return err
}