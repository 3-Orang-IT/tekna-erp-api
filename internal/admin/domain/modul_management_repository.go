package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type ModulManagementRepository interface {
	CreateModul(modul *entity.Modul) error
	GetModuls() ([]entity.Modul, error)
	GetModulByID(id string) (*entity.Modul, error)
	UpdateModul(id string, modul *entity.Modul) error
	DeleteModul(id string) error
}
