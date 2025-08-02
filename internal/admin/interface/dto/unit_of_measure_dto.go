package dto

type CreateUnitOfMeasureInput struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}

type UpdateUnitOfMeasureInput struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}

type UnitOfMeasureResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}