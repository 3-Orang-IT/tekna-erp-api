package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type CityManagementUsecase interface {
	CreateCity(city *entity.City) error
	GetCities(page, limit int, search string) ([]entity.City, error) // Added search parameter
	GetCityByID(id string) (*entity.City, error)
	UpdateCity(id string, city *entity.City) error
	DeleteCity(id string) error
	GetProvinces(page, limit int, search string) ([]entity.Province, error) // Added GetProvinces method
}

type cityManagementUsecase struct {
	repo adminRepository.CityManagementRepository
}

func NewCityManagementUsecase(r adminRepository.CityManagementRepository) CityManagementUsecase {
	return &cityManagementUsecase{repo: r}
}

func (u *cityManagementUsecase) CreateCity(city *entity.City) error {
	return u.repo.CreateCity(city)
}

func (u *cityManagementUsecase) GetCities(page, limit int, search string) ([]entity.City, error) {
	return u.repo.GetCities(page, limit, search)
}

func (u *cityManagementUsecase) GetCityByID(id string) (*entity.City, error) {
	return u.repo.GetCityByID(id)
}

func (u *cityManagementUsecase) UpdateCity(id string, city *entity.City) error {
	return u.repo.UpdateCity(id, city)
}

func (u *cityManagementUsecase) DeleteCity(id string) error {
	return u.repo.DeleteCity(id)
}

func (u *cityManagementUsecase) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
	return u.repo.GetProvinces(page, limit, search) // Assuming the repository has this method
}
