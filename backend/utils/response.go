package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	UnauthorizedCode     = 1001
	ForbiddenCode        = 1002
	NotFoundCode         = 1003
	MethodNotAllowedCode = 1004
	SuccessCode          = 2000
	ErrorCode            = 2001
	ServerErrorCode      = 2002
)

func Success(c *gin.Context, data interface{}, msg string) {
	if msg == "" {
		msg = "success"
	}

	c.JSON(http.StatusOK, Response{
		Code: SuccessCode,
		Msg:  msg,
		Data: data,
	})
}

func Error(c *gin.Context, status int, code int, msg string) {
	c.JSON(status, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ServerError(c *gin.Context, msg string) {
	Error(c, http.StatusInternalServerError, ServerErrorCode, msg)
}

func Unauthorized(c *gin.Context, msg string) {
	Error(c, http.StatusUnauthorized, UnauthorizedCode, msg)
}

func Forbidden(c *gin.Context, msg string) {
	Error(c, http.StatusForbidden, ForbiddenCode, msg)
}

func NotFound(c *gin.Context, msg string) {
	Error(c, http.StatusNotFound, NotFoundCode, msg)
}

func MethodNotAllowed(c *gin.Context, msg string) {
	Error(c, http.StatusMethodNotAllowed, MethodNotAllowedCode, msg)
}

func ValidationError(c *gin.Context, msg string, errors map[string]string) {
	c.JSON(http.StatusBadRequest, Response{
		Code: ErrorCode,
		Msg:  msg,
		Data: gin.H{
			"errors": errors,
		},
	})
}
