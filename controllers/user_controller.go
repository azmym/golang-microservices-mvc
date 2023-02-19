package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"golang-microservices-mvc/services"
	"golang-microservices-mvc/utils"
)

type (
	UserControllerInterface interface {
		GetUser(*gin.Context)
	}
	userController struct {
	}
)

var (
	UserController UserControllerInterface
)

func init() {
	UserController = &userController{}
}

func (uc *userController) GetUser(c *gin.Context) {
	//grab the user_id
	userIdParam := c.Param("user_id")
	//validate the user_id
	userId, err := cast.ToInt64E(userIdParam)
	///
	if err != nil {
		apiErr := utils.ApplicationError{
			Message:    "user_id should be int.",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	user, apiErr := services.UserService.GetUserById(userId)
	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	//Marshal
	c.JSON(http.StatusOK, user)
}
