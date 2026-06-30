package services

import (
	"backend/apperr"
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type ExerciseService struct {
	ExerciseRepository *repositories.ExerciseRepository
}

// NewExerciseService 创建动作业务服务。
func NewExerciseService() *ExerciseService {
	return &ExerciseService{
		ExerciseRepository: repositories.NewExerciseRepository(),
	}
}

// Create 创建当前用户的动作，并校验名称不能为空、同一用户下名称唯一。
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

// GetExercises 查询当前用户的动作列表，支持关键词和肌群筛选。
func (s *ExerciseService) GetExercises(userId int64, req dto.ExerciseSearchRequest) ([]models.Exercise, error) {
	req.Keyword = strings.TrimSpace(req.Keyword)

	return s.ExerciseRepository.FindByUserId(userId, req)
}

// GetExercise 查询当前用户的单个动作详情。
func (s *ExerciseService) GetExercise(userId int64, exerciseId int64) (models.Exercise, error) {
	exercise, err := s.ExerciseRepository.FindByIdAndUserId(exerciseId, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Exercise{}, apperr.ErrExerciseNotFound
	}
	return exercise, err
}

// UpdateExercise 更新当前用户的动作，并校验名称不能为空、同一用户下名称唯一。
func (s *ExerciseService) UpdateExercise(userId, exerciseId int64, req dto.ExerciseUpdateRequest) error {
	req.Name = strings.TrimSpace(req.Name)
	req.Note = strings.TrimSpace(req.Note)
	if req.Name == "" {
		return apperr.ErrInvalidExerciseName
	}

	// 验重，需要排除当前动作
	found, err := s.ExerciseRepository.ExistsByNameAndUserIdExcludeId(req.Name, userId, exerciseId)
	if err != nil {
		return err
	}
	if found {
		return apperr.ErrExerciseNameExisted
	}

	// 创建Exercise
	exercise := &models.Exercise{
		Id:          exerciseId,
		Category:    *req.Category,
		MuscleGroup: *req.MuscleGroup,
		Name:        req.Name,
		Note:        req.Note,
		UserId:      userId,
	}

	if err := s.ExerciseRepository.Update(exercise); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // 23505 是 PostgreSQL 的唯一索引冲突错误代码
			return apperr.ErrExerciseNameExisted
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperr.ErrExerciseNotFound
		}
		return err
	}

	return nil
}

// DeleteExercise 删除当前用户的动作。
func (s *ExerciseService) DeleteExercise(userId, exerciseId int64) error {
	if err := s.ExerciseRepository.Delete(exerciseId, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperr.ErrExerciseNotFound
		}
		return err
	}
	return nil
}
