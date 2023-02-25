package app

import "golang-microservices-mvc/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.UserController.GetUser)
}
