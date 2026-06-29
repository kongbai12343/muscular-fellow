package controllers

import (
	"backend/apperr"
	"backend/dto"
	"backend/services"
	"backend/utils"
	validator "backend/validator"
	"errors"
	"net/http"

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
	default:
		utils.ServerError(c, "服务器内部错误")
	}
}

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

func (ctrl *ExerciseController) GetExercises(c *gin.Context) {
	userId, ok := getUserID(c)
	if !ok {
		utils.Unauthorized(c, "身份验证失败")
		return
	}

	var req dto.ExerciseSearchRequest
	if !validator.BindJSON(c, &req) {
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

func (ctrl *ExerciseController) GetExercise(c *gin.Context) {
}

func (ctrl *ExerciseController) UpdateExercise(c *gin.Context) {
}

func (ctrl *ExerciseController) DeleteExercise(c *gin.Context) {
}
