package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type AreaManagementRepository interface {
	CreateArea(area *entity.Area) error
	GetAreas(page, limit int, search string) ([]entity.Area, error)
	GetAreaByID(id string) (*entity.Area, error)
	UpdateArea(id string, area *entity.Area) error
	DeleteArea(id string) error
	// Method to get total count of areas for pagination
	GetAreasCount(search string) (int64, error)
}
