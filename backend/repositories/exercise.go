package repositories

import (
	"backend/database"
	"backend/models"
)

type ExerciseRepository struct {
}

func NewExerciseRepository() *ExerciseRepository {
	return &ExerciseRepository{}
}

func (r *ExerciseRepository) Create(exercise *models.Exercise) error {
	return database.DB.Create(exercise).Error
}

func (r *ExerciseRepository) FindByUserId(userId int64) ([]models.Exercise, error) {
	var exercises []models.Exercise
	return exercises, database.DB.Where("user_id = ?", userId).Find(&exercises).Error
}

func (r *ExerciseRepository) ExistsByNameAndUserId(name string, userId int64) (bool, error) {
	var count int64
	err := database.DB.Model(&models.Exercise{}).Where("name = ? AND user_id = ?", name, userId).Count(&count).Error
	return count > 0, err
}
