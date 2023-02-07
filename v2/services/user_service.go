package services

import (
	"golang-microservices-mvc/v2/domain"
	"golang-microservices-mvc/v2/utils"
)

func GetUserById(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUserById(userId)
}
