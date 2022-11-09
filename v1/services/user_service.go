package services

import (
	"golang-microservices-mvc/v1/domain"
)

func GetUserById(userId int64) (*domain.User, error) {
	return domain.GetUserById(userId)
}
