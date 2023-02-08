package services

import (
	"golang-microservices-mvc/domain"
	"golang-microservices-mvc/utils"
)

func GetUserById(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUserById(userId)
}
