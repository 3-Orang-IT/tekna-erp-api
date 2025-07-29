package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type TravelCostManagementRepository interface {
	CreateTravelCost(travelCost *entity.TravelCost) error
	GetTravelCosts(page, limit int, search string) ([]entity.TravelCost, error)
	GetTravelCostByID(id string) (*entity.TravelCost, error)
	UpdateTravelCost(id string, travelCost *entity.TravelCost) error
	DeleteTravelCost(id string) error
	// Method to get total count of travel costs for pagination
	GetTravelCostsCount(search string) (int64, error)
	// Method to get the last travel cost for code generation
	GetLastTravelCost() (*entity.TravelCost, error)
}
