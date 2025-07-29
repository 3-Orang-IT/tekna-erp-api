package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type TravelCostManagementUsecase interface {
	CreateTravelCost(travelCost *entity.TravelCost) error
	GetTravelCosts(page, limit int, search string) ([]entity.TravelCost, error)
	GetTravelCostByID(id string) (*entity.TravelCost, error)
	UpdateTravelCost(id string, travelCost *entity.TravelCost) error
	DeleteTravelCost(id string) error
	GetTravelCostsCount(search string) (int64, error) // Method to get total count of travel costs for pagination
	GetLastTravelCost() (*entity.TravelCost, error) // Method to get the last travel cost for code generation
}

type travelCostManagementUsecase struct {
	repo adminRepository.TravelCostManagementRepository
}

func NewTravelCostManagementUsecase(r adminRepository.TravelCostManagementRepository) TravelCostManagementUsecase {
	return &travelCostManagementUsecase{repo: r}
}

func (u *travelCostManagementUsecase) CreateTravelCost(travelCost *entity.TravelCost) error {
	return u.repo.CreateTravelCost(travelCost)
}

func (u *travelCostManagementUsecase) GetTravelCosts(page, limit int, search string) ([]entity.TravelCost, error) {
	return u.repo.GetTravelCosts(page, limit, search)
}

func (u *travelCostManagementUsecase) GetTravelCostByID(id string) (*entity.TravelCost, error) {
	return u.repo.GetTravelCostByID(id)
}

func (u *travelCostManagementUsecase) UpdateTravelCost(id string, travelCost *entity.TravelCost) error {
	return u.repo.UpdateTravelCost(id, travelCost)
}

func (u *travelCostManagementUsecase) DeleteTravelCost(id string) error {
	return u.repo.DeleteTravelCost(id)
}

// GetTravelCostsCount gets the total count of travel costs for pagination
func (u *travelCostManagementUsecase) GetTravelCostsCount(search string) (int64, error) {
	return u.repo.GetTravelCostsCount(search)
}

// GetLastTravelCost gets the last travel cost for code generation
func (u *travelCostManagementUsecase) GetLastTravelCost() (*entity.TravelCost, error) {
	return u.repo.GetLastTravelCost()
}
