package services

import (
	"backend/apperr"
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

type ExerciseService struct {
	ExerciseRepository *repositories.ExerciseRepository
}

func NewExerciseService() *ExerciseService {
	return &ExerciseService{
		ExerciseRepository: repositories.NewExerciseRepository(),
	}
}

func (s *ExerciseService) Create(userId int64, req dto.ExerciseRequest) error {
	req.Name = strings.TrimSpace(req.Name)
	req.Note = strings.TrimSpace(req.Note)
	if req.Name == "" {
		return apperr.ErrInvalidExerciseName
	}
	// 验重
	found, err := s.ExerciseRepository.ExistsByNameAndUserId(req.Name, userId)
	if err != nil {
		return err
	}
	if found {
		return apperr.ErrExerciseNameExisted
	}

	// 创建Exercise
	exercise := &models.Exercise{
		Category:    *req.Category,
		MuscleGroup: *req.MuscleGroup,
		Name:        req.Name,
		Note:        req.Note,
		UserId:      userId,
	}
	if err := s.ExerciseRepository.Create(exercise); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // 23505 是 PostgreSQL 的唯一索引冲突错误代码
			return apperr.ErrExerciseNameExisted
		}
		return err
	}
	return nil
}

func (s *ExerciseService) GetExercises(userId int64, req dto.ExerciseSearchRequest) ([]models.Exercise, error) {
	req.Name = strings.TrimSpace(req.Name)
	return s.ExerciseRepository.FindByUserId(userId)
}
