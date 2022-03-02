package repositories

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreateUserMock(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	repo := NewPostgreUserRepository(db)

	query := regexp.QuoteMeta("INSERT INTO users (name, email, phone) VALUES ($1, $2, $3)  RETURNING id")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(query).WithArgs(userTest.Name, userTest.Email, userTest.Phone).WillReturnRows(rows)

	id, err := repo.CreateUser(userTest)

	assert.NotNil(t, id)
	assert.NoError(t, err)
}

func TestGetUserByIdMock(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	repo := NewPostgreUserRepository(db)

	query := regexp.QuoteMeta("SELECT * FROM users WHERE id=$1")
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(1, userTest.Name, userTest.Email, userTest.Phone)
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	user, err := repo.GetUserById(1)

	assert.NotNil(t, user)
	assert.NoError(t, err)
}


func TestCheckUserExistsMock(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	repo := NewPostgreUserRepository(db)

	query := regexp.QuoteMeta("SELECT id FROM users WHERE email=$1 OR phone=$2")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(query).WithArgs(userTest.Email, userTest.Phone).WillReturnRows(rows)

	ok, err := repo.CheckUserExists(userTest)

	assert.True(t, ok)
	assert.NoError(t, err)
}
