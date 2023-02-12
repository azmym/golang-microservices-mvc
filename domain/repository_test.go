package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindByIdIUserNotExist(t *testing.T) {
	found, user := users.findById(12)

	assert.Nil(t, user)
	assert.EqualValues(t, false, found)
}

func TestFindByIdIUserExist(t *testing.T) {
	found, user := users.findById(123)

	assert.NotNil(t, user)
	assert.EqualValues(t, true, found)

	assert.EqualValues(t, 123, user.Id)
}

func TestGetUserByIdNotExist(t *testing.T) {
	repository := UserRepository{}
	user, applicationError := repository.FindUserById(0)

	assert.Nil(t, user)
	assert.NotNil(t, applicationError)

	assert.EqualValues(t, applicationError.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, "not_found", applicationError.Code)
	assert.EqualValues(t, "there is no user with id= 0", applicationError.Message)
}

func TestGetUserById(t *testing.T) {
	repository := UserRepository{}
	user, applicationError := repository.FindUserById(123)

	assert.Nil(t, applicationError)
	assert.NotNil(t, user)

	assert.EqualValues(t, user.Id, 123)
	assert.EqualValues(t, "Hello", user.FirstName)
	assert.EqualValues(t, "world", user.LastName)
}
