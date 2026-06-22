package controllers

import (
	"backend/apperr"
	"backend/dto"
	"backend/services"
	"backend/utils"
	validator "backend/validator"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func handleUserError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, apperr.ErrInvalidCredentials):
		utils.Error(c, http.StatusUnauthorized, utils.ErrorCode, err.Error())
	case errors.Is(err, apperr.ErrInvalidUsername):
		utils.Error(c, http.StatusBadRequest, utils.ErrorCode, err.Error())
	case errors.Is(err, apperr.ErrEmailExists),
		errors.Is(err, apperr.ErrUsernameExists),
		errors.Is(err, apperr.ErrUserAlreadyExists):
		utils.Error(c, http.StatusConflict, utils.ErrorCode, err.Error())
	default:
		utils.ServerError(c, "服务器内部错误")
	}
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (ctrl *UserController) Login(c *gin.Context) {
	var req dto.UserLogin
	if !validator.BindJSON(c, &req) {
		return
	}

	user, err := ctrl.userService.Login(req)
	if err != nil {
		handleUserError(c, err)
		return
	}

	token, err := utils.GenerateToken(user.Id, user.Username)
	if err != nil {
		utils.ServerError(c, "服务器内部错误")
		return
	}

	resp := &dto.UserResponse{
		Token: token,
		UserInfo: dto.User{
			ID:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	utils.Success(c, resp, "登录成功")
}

func (ctrl *UserController) Register(c *gin.Context) {
	var req dto.UserRegister
	if !validator.BindJSON(c, &req) {
		return
	}

	err := ctrl.userService.Register(req)
	if err != nil {
		handleUserError(c, err)
		return
	}

	utils.Success(c, nil, "注册成功")
}
