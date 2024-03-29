package services

import (
	"golang-microservices-mvc/domain"
	"golang-microservices-mvc/utils"
)

type (
	UserServiceInterface interface {
		GetUserById(int64) (*domain.User, *utils.ApplicationError)
	}
	userService struct {
	}
)

var (
	UserService UserServiceInterface
)

func init() {
	UserService = &userService{}
}

func (us *userService) GetUserById(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserRepository.FindUserById(userId)
}
