package controllers

import (
	"backend/apperr"
	"backend/dto"
	"backend/services"
	"backend/utils"
	validator "backend/validator"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExerciseController struct {
	exerciseService *services.ExerciseService
}

func NewExerciseController() *ExerciseController {
	return &ExerciseController{
		exerciseService: services.NewExerciseService(),
	}
}

func getUserID(c *gin.Context) (int64, bool) {
	userId, exists := c.Get("userId")
	if !exists {
		return 0, false
	}

	userIdInt, ok := userId.(int64)
	if !ok || userIdInt <= 0 {
		return 0, false
	}

	return userIdInt, true
}

func handleExerciseError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, apperr.ErrInvalidExerciseName):
		utils.Error(c, http.StatusBadRequest, utils.ErrorCode, err.Error())
	case errors.Is(err, apperr.ErrExerciseNameExisted):
		utils.Error(c, http.StatusConflict, utils.ErrorCode, err.Error())
	case errors.Is(err, apperr.ErrExerciseNotFound):
		utils.NotFound(c, err.Error())
	default:
		utils.ServerError(c, "服务器内部错误")
	}
}

// 创建动作
func (ctrl *ExerciseController) Create(c *gin.Context) {
	var req dto.ExerciseRequest
	userId, ok := getUserID(c)
	if !ok {
		utils.Unauthorized(c, "身份验证失败")
		return
	}

	if !validator.BindJSON(c, &req) {
		return
	}

	err := ctrl.exerciseService.Create(userId, req)
	if err != nil {
		handleExerciseError(c, err)
		return
	}
	utils.Success(c, nil, "创建成功")
}

// 获取所有动作
func (ctrl *ExerciseController) GetExercises(c *gin.Context) {
	userId, ok := getUserID(c)
	if !ok {
		utils.Unauthorized(c, "身份验证失败")
		return
	}

	var req dto.ExerciseSearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, utils.ErrorCode, "参数错误")
		return
	}

	exercises, err := ctrl.exerciseService.GetExercises(userId, req)
	if err != nil {
		handleExerciseError(c, err)
		return
	}

	resp := make([]dto.ExerciseResponse, 0, len(exercises))
	for _, exercise := range exercises {
		resp = append(resp, dto.ExerciseResponse{
			ID:          exercise.Id,
			Name:        exercise.Name,
			MuscleGroup: exercise.MuscleGroup,
			Category:    exercise.Category,
			Note:        exercise.Note,
			CreatedAt:   utils.FormatDateTime(exercise.CreatedAt),
		})
	}
	utils.Success(c, resp, "success")
}

// 获取动作详情
func (ctrl *ExerciseController) GetExercise(c *gin.Context) {
	userId, ok := getUserID(c)
	if !ok {
		utils.Unauthorized(c, "身份验证失败")
		return
	}

	exerciseIdStr := c.Param("id")

	if exerciseIdStr == "" {
		utils.Error(c, http.StatusBadRequest, utils.ErrorCode, "参数错误")
		return
	}

	exerciseId, err := strconv.ParseInt(exerciseIdStr, 10, 64)
	if err != nil || exerciseId <= 0 {
		utils.Error(c, http.StatusBadRequest, utils.ErrorCode, "参数错误")
		return
	}

	exercise, err := ctrl.exerciseService.GetExercise(userId, exerciseId)
	if err != nil {
		handleExerciseError(c, err)
		return
	}
	utils.Success(c, dto.ExerciseResponse{
		ID:          exercise.Id,
		Name:        exercise.Name,
		MuscleGroup: exercise.MuscleGroup,
		Category:    exercise.Category,
		Note:        exercise.Note,
		CreatedAt:   utils.FormatDateTime(exercise.CreatedAt),
	}, "success")
}

// 编辑动作
func (ctrl *ExerciseController) UpdateExercise(c *gin.Context) {
	var req dto.ExerciseUpdateRequest
	userId, ok := getUserID(c)
	if !ok {
		utils.Unauthorized(c, "身份验证失败")
		return
	}

	exerciseIdStr := c.Param("id")

	if exerciseIdStr == "" {
		utils.Error(c, http.StatusBadRequest, utils.ErrorCode, "参数错误")
		return
	}

	exerciseId, err := strconv.ParseInt(exerciseIdStr, 10, 64)
	if err != nil || exerciseId <= 0 {
		utils.Error(c, http.StatusBadRequest, utils.ErrorCode, "参数错误")
		return
	}

	if !validator.BindJSON(c, &req) {
		return
	}
	if err := ctrl.exerciseService.UpdateExercise(userId, exerciseId, req); err != nil {
		handleExerciseError(c, err)
		return
	}
	utils.Success(c, nil, "更新成功")
}

// 删除动作
func (ctrl *ExerciseController) DeleteExercise(c *gin.Context) {
	userId, ok := getUserID(c)
	if !ok {
		utils.Unauthorized(c, "身份验证失败")
		return
	}

	exerciseIdStr := c.Param("id")

	if exerciseIdStr == "" {
		utils.Error(c, http.StatusBadRequest, utils.ErrorCode, "参数错误")
		return
	}

	exerciseId, err := strconv.ParseInt(exerciseIdStr, 10, 64)
	if err != nil || exerciseId <= 0 {
		utils.Error(c, http.StatusBadRequest, utils.ErrorCode, "参数错误")
		return
	}

	if err := ctrl.exerciseService.DeleteExercise(userId, exerciseId); err != nil {
		handleExerciseError(c, err)
		return
	}
	utils.Success(c, nil, "删除成功")
}
