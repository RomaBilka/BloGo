package repositories

import (
	"testing"

	"github.com/RomaBiliak/BloGo/project-layout/tests/internal/models"
	"github.com/stretchr/testify/assert"
)

func createUser(t *testing.T) models.UserId {
	id, err := testUserRepository.CreateUser(userTest)
	assert.NotNil(t, id)
	assert.NoError(t, err)
	return id
}

func TestCreateUser(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()
	createUser(t)
}

func TestGetUserById(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()
	id := createUser(t)

	u, err := testUserRepository.GetUserById(id)
	assert.Equal(t, u.Id, id)
	assert.NoError(t, err)
}

func TestCheckUserExistsTrue(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()
	createUser(t)

	ok, err := testUserRepository.CheckUserExists(userTest)

	assert.True(t, ok)
	assert.NoError(t, err)
}