package dto

type CreateDivisionInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateDivisionInput struct {
	Name string `json:"name" binding:"required"`
}
