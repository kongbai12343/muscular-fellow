package services

import "backend/repositories"

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository: repositories.NewUserRepository(),
	}
}
