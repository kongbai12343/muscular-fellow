package services

import (
	"backend/apperr"
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository: repositories.NewUserRepository(),
	}
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func normalizeUsername(username string) string {
	return strings.TrimSpace(username)
}

func (s *UserService) Login(req dto.UserLogin) (*models.User, error) {
	user, found, err := s.UserRepository.FindByEmail(normalizeEmail(req.Email))
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, apperr.ErrInvalidCredentials
	}

	if !utils.VerifyPassword(user.Password, req.Password) {
		return nil, apperr.ErrInvalidCredentials
	}

	return user, nil
}

func (s *UserService) Register(req dto.UserRegister) error {
	email := normalizeEmail(req.Email)
	username := normalizeUsername(req.Username)
	if len(username) < 2 {
		return apperr.ErrInvalidUsername
	}

	user, found, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return err
	}
	if found {
		return apperr.ErrEmailExists
	}

	user, found, err = s.UserRepository.FindByUsername(username)
	if err != nil {
		return err
	}
	if found {
		return apperr.ErrUsernameExists
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user = &models.User{
		Email:    email,
		Password: hashedPassword,
		Username: username,
	}

	err = s.UserRepository.Create(user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // 23505 是 PostgreSQL 的唯一索引冲突错误代码
			return apperr.ErrUserAlreadyExists
		}
		return err
	}

	return nil
}
