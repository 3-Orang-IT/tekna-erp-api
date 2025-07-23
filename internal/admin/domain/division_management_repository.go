package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type DivisionManagementRepository interface {
	CreateDivision(division *entity.Division) error
	GetDivisions(page, limit int) ([]entity.Division, error)
	GetDivisionByID(id string) (*entity.Division, error)
	UpdateDivision(id string, division *entity.Division) error
	DeleteDivision(id string) error
}
