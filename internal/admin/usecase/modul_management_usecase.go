package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type ModulManagementUsecase interface {
	CreateModul(modul *entity.Modul) error
	GetModuls() ([]entity.Modul, error)
	GetModulByID(id string) (*entity.Modul, error)
	UpdateModul(id string, modul *entity.Modul) error
	DeleteModul(id string) error
}

type modulManagementUsecase struct {
	repo adminRepository.ModulManagementRepository
}

func NewModulManagementUsecase(r adminRepository.ModulManagementRepository) ModulManagementUsecase {
	return &modulManagementUsecase{repo: r}
}

func (u *modulManagementUsecase) CreateModul(modul *entity.Modul) error {
	return u.repo.CreateModul(modul)
}

func (u *modulManagementUsecase) GetModuls() ([]entity.Modul, error) {
	return u.repo.GetModuls()
}

func (u *modulManagementUsecase) GetModulByID(id string) (*entity.Modul, error) {
	return u.repo.GetModulByID(id)
}

func (u *modulManagementUsecase) UpdateModul(id string, modul *entity.Modul) error {
	return u.repo.UpdateModul(id, modul)
}

func (u *modulManagementUsecase) DeleteModul(id string) error {
	return u.repo.DeleteModul(id)
}
