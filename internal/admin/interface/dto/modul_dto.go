package dto

type CreateModulInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateModulInput struct {
	Name string `json:"name"`
}
