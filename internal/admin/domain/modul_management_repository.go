package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type ModulManagementRepository interface {
	CreateModul(modul *entity.Modul) error
	GetModuls(page, limit int) ([]entity.Modul, error) // Updated to include pagination
	GetModulByID(id string) (*entity.Modul, error)
	UpdateModul(id string, modul *entity.Modul) error
	DeleteModul(id string) error
}
