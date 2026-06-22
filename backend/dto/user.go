package dto

type UserLogin struct {
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

type UserRegister struct {
	Username string `json:"username" binding:"required,min=2,max=100"`
	Password string `json:"password" binding:"required,min=8,max=72"`
	Email    string `json:"email" binding:"required,email,max=100"`
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserResponse struct {
	UserInfo User   `json:"user"`
	Token    string `json:"token"`
}
