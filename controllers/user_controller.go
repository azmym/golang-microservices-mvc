package controllers

import (
	"encoding/json"
	"golang-microservices-mvc/services"
	"golang-microservices-mvc/utils"
	"net/http"
	"strconv"
)

type (
	userControllerInterface interface {
		GetUser(http.ResponseWriter, *http.Request)
	}
	userController struct {
	}
)

var (
	UserController userController
)

func (uc *userController) GetUser(res http.ResponseWriter, req *http.Request) {
	//grab the user_id
	userIdParam := req.URL.Query().Get("user_id")
	//validate the user_id
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := utils.ApplicationError{
			Message:    "user_id should be int.",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		errorResp, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(errorResp)
		return
	}
	user, apiErr := services.UserService.GetUserById(userId)
	if apiErr != nil {
		errorResp, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(errorResp)
		return
	}
	//Marshal
	userDto, _ := json.Marshal(user)
	res.WriteHeader(http.StatusOK)
	res.Write(userDto)
}
