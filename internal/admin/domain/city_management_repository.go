package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type CityManagementRepository interface {
	CreateCity(city *entity.City) error
	GetCities(page, limit int, search string) ([]entity.City, error) // Added search parameter
	GetCitiesCount(search string) (int64, error) // Added count method for pagination
	GetCityByID(id string) (*entity.City, error)
	UpdateCity(id string, city *entity.City) error
	DeleteCity(id string) error
	GetProvinces(page, limit int, search string) ([]entity.Province, error) // Added GetProvinces method
}
