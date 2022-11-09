package controllers

import (
	"encoding/json"
	"golang-microservices-mvc/v1/services"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	//grab the user_id
	userIdParam := req.URL.Query().Get("user_id")
	//validate the user_id
	userId, error := strconv.ParseInt(userIdParam, 10, 64)
	if error != nil {
		res.WriteHeader(404)
		res.Write([]byte("user_id should be int."))
		return
	}
	user, err := services.GetUserById(userId)
	if err != nil {
		res.WriteHeader(200)
		res.Write([]byte(err.Error()))
		return
	}
	//Marshal
	userDto, _ := json.Marshal(user)
	res.Write(userDto)
}
