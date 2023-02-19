package app

import (
	"net/http"

	"golang-microservices-mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.UserController.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
