package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type AreaManagementUsecase interface {
	CreateArea(area *entity.Area) error
	GetAreas(page, limit int, search string) ([]entity.Area, error)
	GetAreaByID(id string) (*entity.Area, error)
	UpdateArea(id string, area *entity.Area) error
	DeleteArea(id string) error
	GetAreasCount(search string) (int64, error) // Method to get total count of areas for pagination
}

type areaManagementUsecase struct {
	repo adminRepository.AreaManagementRepository
}

func NewAreaManagementUsecase(r adminRepository.AreaManagementRepository) AreaManagementUsecase {
	return &areaManagementUsecase{repo: r}
}

func (u *areaManagementUsecase) CreateArea(area *entity.Area) error {
	return u.repo.CreateArea(area)
}

func (u *areaManagementUsecase) GetAreas(page, limit int, search string) ([]entity.Area, error) {
	return u.repo.GetAreas(page, limit, search)
}

func (u *areaManagementUsecase) GetAreaByID(id string) (*entity.Area, error) {
	return u.repo.GetAreaByID(id)
}

func (u *areaManagementUsecase) UpdateArea(id string, area *entity.Area) error {
	return u.repo.UpdateArea(id, area)
}

func (u *areaManagementUsecase) DeleteArea(id string) error {
	return u.repo.DeleteArea(id)
}

// GetAreasCount gets the total count of areas for pagination
func (u *areaManagementUsecase) GetAreasCount(search string) (int64, error) {
	return u.repo.GetAreasCount(search)
}
