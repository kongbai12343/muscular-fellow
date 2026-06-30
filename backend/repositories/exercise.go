package repositories

import (
	"backend/database"
	"backend/dto"
	"backend/models"

	"gorm.io/gorm"
)

type ExerciseRepository struct {
}

// NewExerciseRepository 创建动作数据仓库。
func NewExerciseRepository() *ExerciseRepository {
	return &ExerciseRepository{}
}

// Create 新增动作记录。
func (r *ExerciseRepository) Create(exercise *models.Exercise) error {
	return database.DB.Create(exercise).Error
}

// FindByUserId 查询当前用户的动作列表，支持关键词和肌群筛选。
func (r *ExerciseRepository) FindByUserId(userId int64, req dto.ExerciseSearchRequest) ([]models.Exercise, error) {
	var exercises []models.Exercise
	var query = database.DB.Model(&models.Exercise{}).Where("user_id = ?", userId)
	if req.MuscleGroup != nil {
		query = query.Where("muscle_group = ?", *req.MuscleGroup)
	}
	if req.Keyword != "" {
		query = query.Where("name LIKE ?", "%"+req.Keyword+"%")
	}
	return exercises, query.Order("created_at DESC, id DESC").Find(&exercises).Error
}

// ExistsByNameAndUserId 判断当前用户下是否已存在同名动作。
func (r *ExerciseRepository) ExistsByNameAndUserId(name string, userId int64) (bool, error) {
	var count int64
	err := database.DB.Model(&models.Exercise{}).Where("name = ? AND user_id = ?", name, userId).Count(&count).Error
	return count > 0, err
}

// ExistsByNameAndUserIdExcludeId 判断当前用户下是否存在除指定动作外的同名动作。
func (r *ExerciseRepository) ExistsByNameAndUserIdExcludeId(name string, userId, exerciseId int64) (bool, error) {
	var count int64
	err := database.DB.Model(&models.Exercise{}).Where("name = ? AND user_id = ? AND id <> ?", name, userId, exerciseId).Count(&count).Error
	return count > 0, err
}

// FindByIdAndUserId 按动作 ID 和用户 ID 查询动作，避免越权访问。
func (r *ExerciseRepository) FindByIdAndUserId(id int64, userId int64) (models.Exercise, error) {
	var exercise models.Exercise
	return exercise, database.DB.Model(&models.Exercise{}).Where("id = ? AND user_id = ?", id, userId).First(&exercise).Error
}

// Update 按动作 ID 和用户 ID 更新动作，未命中记录时返回 gorm.ErrRecordNotFound。
func (r *ExerciseRepository) Update(exercise *models.Exercise) error {
	result := database.DB.Model(&models.Exercise{}).
		Where("id = ? AND user_id = ?", exercise.Id, exercise.UserId).
		Updates(map[string]interface{}{
			"category":     exercise.Category,
			"muscle_group": exercise.MuscleGroup,
			"name":         exercise.Name,
			"note":         exercise.Note,
		})

	if result.Error != nil {
		return result.Error
	}
	// Updates 没有影响到任何数据行
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Delete 按动作 ID 和用户 ID 删除动作，未命中记录时返回 gorm.ErrRecordNotFound。
func (r *ExerciseRepository) Delete(id int64, userId int64) error {
	result := database.DB.Model(&models.Exercise{}).Where("id = ? AND user_id = ?", id, userId).Delete(&models.Exercise{})
	if result.Error != nil {
		return result.Error
	}
	// Delete 没有影响到任何数据行
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
