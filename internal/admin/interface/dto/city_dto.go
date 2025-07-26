package dto

type CreateCityInput struct {
	Name       string `json:"name" binding:"required"`
	ProvinceID uint   `json:"province_id" binding:"required"`
}

type UpdateCityInput struct {
	Name       string `json:"name"`
	ProvinceID uint   `json:"province_id"`
}

type CityResponse struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	Province string          `json:"province"`
}
