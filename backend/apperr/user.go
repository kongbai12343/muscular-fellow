package apperr

import "errors"

var (
	ErrInvalidCredentials = errors.New("邮箱或密码错误")
	ErrEmailExists        = errors.New("邮箱已存在")
	ErrUsernameExists     = errors.New("用户名已存在")
	ErrUserAlreadyExists  = errors.New("邮箱或用户名已存在")
	ErrInvalidUsername    = errors.New("用户名长度不能小于 2")
)
