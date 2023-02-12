package services

import (
	"golang-microservices-mvc/domain"
	"golang-microservices-mvc/utils"
)

type (
	UserServiceInterface interface {
		GetUserById(int64) (*domain.User, *utils.ApplicationError)
	}
	UserService struct {
		UserRepository domain.UserRepositoryInterface
	}
)

func (us UserService) GetUserById(userId int64) (*domain.User, *utils.ApplicationError) {
	return us.UserRepository.FindUserById(userId)
}
