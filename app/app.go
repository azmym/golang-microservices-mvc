package app

import (
	"golang-microservices-mvc/controllers"
	"golang-microservices-mvc/domain"
	"golang-microservices-mvc/services"
	"net/http"
)

func StartApp() {
	userController := di()
	http.HandleFunc("/users", userController.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func di() *controllers.UserController {
	return &controllers.UserController{
		UserService: &services.UserService{
			UserRepository: &domain.UserRepository{},
		},
	}
}
