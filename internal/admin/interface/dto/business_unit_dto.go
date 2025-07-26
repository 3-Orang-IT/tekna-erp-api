package dto

type CreateBusinessUnitInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateBusinessUnitInput struct {
	Name string `json:"name" binding:"required"`
}
