package apperr

import "errors"

var (
	ErrExerciseNameExisted = errors.New("动作已存在")
	ErrInvalidExerciseName = errors.New("动作名称不能为空")
	ErrExerciseNotFound    = errors.New("动作不存在")
)
