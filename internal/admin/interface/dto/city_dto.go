package dto

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

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
	Province entity.Province `json:"province"`
}
