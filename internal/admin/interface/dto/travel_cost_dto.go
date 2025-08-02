package dto

type CreateTravelCostInput struct {
	Name  string  `json:"name" binding:"required"`
	Unit  string  `json:"unit" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type UpdateTravelCostInput struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Unit  string  `json:"unit"`
	Price float64 `json:"price"`
}

type TravelCostResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Unit      string  `json:"unit"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
