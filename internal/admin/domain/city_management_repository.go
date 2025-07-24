package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type CityManagementRepository interface {
	CreateCity(city *entity.City) error
	GetCities(page, limit int) ([]entity.City, error)
	GetCityByID(id string) (*entity.City, error)
	UpdateCity(id string, city *entity.City) error
	DeleteCity(id string) error
}
