package dto

type CreateDivisionInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateDivisionInput struct {
	Name string `json:"name" binding:"required"`
}

type DivisionResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
