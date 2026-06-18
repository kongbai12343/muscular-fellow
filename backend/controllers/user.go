package controllers

import (
	"backend/dto"
	"backend/services"
	validator "backend/validator"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func Login(c *gin.Context) {
	var req dto.UserLogin
	if !validator.BindJSON(c, &req) {
		return
	}
}

func Register(c *gin.Context) {
	var req dto.UserRegister
	if !validator.BindJSON(c, &req) {
		return
	}
}

func GetUserInfo(c *gin.Context) {
}
