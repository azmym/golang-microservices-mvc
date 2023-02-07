package app

import (
	"golang-microservices-mvc/v2/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
