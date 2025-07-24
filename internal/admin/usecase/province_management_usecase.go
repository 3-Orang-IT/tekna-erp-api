package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type ProvinceManagementUsecase interface {
	CreateProvince(province *entity.Province) error
	GetProvinces(page, limit int) ([]entity.Province, error)
	GetProvinceByID(id string) (*entity.Province, error)
	UpdateProvince(id string, province *entity.Province) error
	DeleteProvince(id string) error
}

type provinceManagementUsecase struct {
	repo adminRepository.ProvinceManagementRepository
}

func NewProvinceManagementUsecase(r adminRepository.ProvinceManagementRepository) ProvinceManagementUsecase {
	return &provinceManagementUsecase{repo: r}
}

func (u *provinceManagementUsecase) CreateProvince(province *entity.Province) error {
	return u.repo.CreateProvince(province)
}

func (u *provinceManagementUsecase) GetProvinces(page, limit int) ([]entity.Province, error) {
	return u.repo.GetProvinces(page, limit)
}

func (u *provinceManagementUsecase) GetProvinceByID(id string) (*entity.Province, error) {
	return u.repo.GetProvinceByID(id)
}

func (u *provinceManagementUsecase) UpdateProvince(id string, province *entity.Province) error {
	return u.repo.UpdateProvince(id, province)
}

func (u *provinceManagementUsecase) DeleteProvince(id string) error {
	return u.repo.DeleteProvince(id)
}
