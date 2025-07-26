package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type BusinessUnitManagementRepository interface {
	CreateBusinessUnit(businessUnit *entity.BusinessUnit) error
	GetBusinessUnits(page, limit int, search string) ([]entity.BusinessUnit, error)
	GetBusinessUnitByID(id string) (*entity.BusinessUnit, error)
	UpdateBusinessUnit(id string, businessUnit *entity.BusinessUnit) error
	DeleteBusinessUnit(id string) error
}
