package services

import (
	"golang-microservices-mvc/domain"
	"golang-microservices-mvc/utils"
)

var (
	UserService userServiceInterface
)

type (
	userServiceInterface interface {
		GetUserById(int64) (*domain.User, *utils.ApplicationError)
	}
	userService struct{}
)

func (us *userService) GetUserById(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserRepository.GetUserById(userId)
}
