package services

import (
	"golang-microservices-mvc/domain"
	"golang-microservices-mvc/utils"
)

type (
	UserRepositoryMock struct{}
)

var (
	FindUserByIdFunc func(userId int64) (*domain.User, *utils.ApplicationError)
)

func (*UserRepositoryMock) FindUserById(userId int64) (*domain.User, *utils.ApplicationError) {
	return FindUserByIdFunc(userId)
}
