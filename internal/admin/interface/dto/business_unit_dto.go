package dto

type CreateBusinessUnitInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateBusinessUnitInput struct {
	Name string `json:"name" binding:"required"`
}

type BusinessUnitResponse struct {
	ID        uint `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}