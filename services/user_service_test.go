package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-microservices-mvc/domain"
	"golang-microservices-mvc/utils"
	"net/http"
	"testing"
)

type (
	userRepositoryMock struct{}
)

var (
	findUserByIdFunc func(userId int64) (*domain.User, *utils.ApplicationError)
)

func (*userRepositoryMock) FindUserById(userId int64) (*domain.User, *utils.ApplicationError) {
	return findUserByIdFunc(userId)
}

func init() {
	domain.UserRepository = &userRepositoryMock{}
}
func TestGetUserByIdIfUserNotExist(test *testing.T) {
	//given
	findUserByIdFunc = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("there is no user with id= %v", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	//when
	user, applicationError := UserService.GetUserById(1)

	//then
	assert.Nil(test, user)
	assert.EqualValues(test, "there is no user with id= 1", applicationError.Message)
	assert.EqualValues(test, http.StatusNotFound, applicationError.StatusCode)
	assert.EqualValues(test, "not_found", applicationError.Code)
}

func TestGetUserByIdIfUserExist(test *testing.T) {
	//given
	findUserByIdFunc = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id:        1,
			FirstName: "Test",
			LastName:  "Golang",
			Email:     "test@golang.com",
		}, nil
	}
	//when
	user, applicationError := UserService.GetUserById(1)

	//then
	assert.Nil(test, applicationError)
	assert.EqualValues(test, 1, user.Id)
	assert.EqualValues(test, "Test", user.FirstName)
	assert.EqualValues(test, "Golang", user.LastName)
	assert.EqualValues(test, "test@golang.com", user.Email)
}
