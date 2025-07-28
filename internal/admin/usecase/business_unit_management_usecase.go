package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type BusinessUnitManagementUsecase interface {
	CreateBusinessUnit(businessUnit *entity.BusinessUnit) error
	GetBusinessUnits(page, limit int, search string) ([]entity.BusinessUnit, error)
	GetBusinessUnitByID(id string) (*entity.BusinessUnit, error)
	UpdateBusinessUnit(id string, businessUnit *entity.BusinessUnit) error
	DeleteBusinessUnit(id string) error
	GetBusinessUnitsCount(search string) (int64, error) // Method to get total count of business units for pagination
}

type businessUnitManagementUsecase struct {
	repo adminRepository.BusinessUnitManagementRepository
}

func NewBusinessUnitManagementUsecase(r adminRepository.BusinessUnitManagementRepository) BusinessUnitManagementUsecase {
	return &businessUnitManagementUsecase{repo: r}
}

func (u *businessUnitManagementUsecase) CreateBusinessUnit(businessUnit *entity.BusinessUnit) error {
	return u.repo.CreateBusinessUnit(businessUnit)
}

func (u *businessUnitManagementUsecase) GetBusinessUnits(page, limit int, search string) ([]entity.BusinessUnit, error) {
	return u.repo.GetBusinessUnits(page, limit, search)
}

func (u *businessUnitManagementUsecase) GetBusinessUnitByID(id string) (*entity.BusinessUnit, error) {
	return u.repo.GetBusinessUnitByID(id)
}

func (u *businessUnitManagementUsecase) UpdateBusinessUnit(id string, businessUnit *entity.BusinessUnit) error {
	return u.repo.UpdateBusinessUnit(id, businessUnit)
}

func (u *businessUnitManagementUsecase) DeleteBusinessUnit(id string) error {
	return u.repo.DeleteBusinessUnit(id)
}

// GetBusinessUnitsCount gets the total count of business units for pagination
func (u *businessUnitManagementUsecase) GetBusinessUnitsCount(search string) (int64, error) {
	return u.repo.GetBusinessUnitsCount(search)
}
