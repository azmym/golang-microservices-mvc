package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/cast"

	"golang-microservices-mvc/services"
	"golang-microservices-mvc/utils"
)

type (
	UserControllerInterface interface {
		GetUser(http.ResponseWriter, *http.Request)
	}
	userController struct {
	}
)

var (
	UserController UserControllerInterface = &userController{}
)

func (uc *userController) GetUser(res http.ResponseWriter, req *http.Request) {
	//grab the user_id
	userIdParam := req.URL.Query().Get("user_id")
	//validate the user_id
	userId, err := cast.ToInt64E(userIdParam)
	///
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
