package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type UnitOfMeasureManagementRepository interface {
	CreateUnitOfMeasure(unitOfMeasure *entity.UnitOfMeasure) error
	GetUnitOfMeasures(page, limit int, search string) ([]entity.UnitOfMeasure, error)
	GetUnitOfMeasureByID(id string) (*entity.UnitOfMeasure, error)
	UpdateUnitOfMeasure(id string, unitOfMeasure *entity.UnitOfMeasure) error
	DeleteUnitOfMeasure(id string) error
	GetUnitOfMeasuresCount(search string) (int64, error) // Method to get total count of units of measure for pagination
}
