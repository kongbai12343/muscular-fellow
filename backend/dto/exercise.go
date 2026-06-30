package dto

type ExerciseRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=50"`
	MuscleGroup *int16 `json:"muscle_group" binding:"required,min=0,max=7"` // 改成 指针，避免0被required 当作空值
	Category    *int16 `json:"category" binding:"required,min=0,max=4"`
	Note        string `json:"note" binding:"max=300"`
}

type ExerciseResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	MuscleGroup int16  `json:"muscle_group"`
	Category    int16  `json:"category"`
	Note        string `json:"note"`
	CreatedAt   string `json:"created_at"`
}

type ExerciseUpdateRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=50"`
	MuscleGroup *int16 `json:"muscle_group" binding:"required,min=0,max=7"`
	Category    *int16 `json:"category" binding:"required,min=0,max=4"`
	Note        string `json:"note" binding:"max=300"`
}

type ExerciseDeleteRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type ExerciseSearchRequest struct {
	Keyword     string `form:"keyword" binding:"omitempty,min=1,max=50"`
	MuscleGroup *int16 `form:"muscle_group" binding:"omitempty,min=0,max=7"`
}
