package handlers

import (
	"database/sql"
	"os"
	"testing"
	"errors"

	"github.com/RomaBilka/BloGo/tests/pkg/database/postgres"
	"github.com/RomaBilka/BloGo/tests/internal/models"
	"github.com/RomaBilka/BloGo/tests/internal/repositories"
	"github.com/RomaBilka/BloGo/tests/internal/services"
	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var testUserRepository *repositories.UserRepository

var uHttp *userHTTP

type errorResponse struct {
	Error string `json:"error"`
}

var db *sql.DB

var userTest createUserRequest

type testData struct {
	Name  string `faker:"name"`
	Email string `faker:"email"`
	Phone string `faker:"e_164_phone_number"`
}

var test = testData{}

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	err = faker.FakeData(&test)
	if err != nil {
		panic(err)
	}

	userTest = createUserRequest{test.Name, test.Email, test.Phone}

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
	testUserRepository = repositories.NewPostgreUserRepository(db)

	userService := services.NewUserService(testUserRepository)
	uHttp = NewUserHttp(userService)

	os.Exit(m.Run())
}

func truncateUsers() error {
	_, err := db.Query("TRUNCATE users CASCADE")
	return err
}

func createTestUser(t *testing.T) models.UserId {
	userId, err := testUserRepository.CreateUser(models.User{Name: userTest.Name, Email: userTest.Email, Phone: userTest.Phone})
	assert.NoError(t, err)
	return userId
}
