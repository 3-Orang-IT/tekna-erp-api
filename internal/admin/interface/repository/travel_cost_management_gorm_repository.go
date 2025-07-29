package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type travelCostManagementRepo struct {
	db *gorm.DB
}

func NewTravelCostManagementRepository(db *gorm.DB) adminRepository.TravelCostManagementRepository {
	return &travelCostManagementRepo{db: db}
}

func (r *travelCostManagementRepo) CreateTravelCost(travelCost *entity.TravelCost) error {
	return r.db.Create(travelCost).Error
}

func (r *travelCostManagementRepo) GetTravelCosts(page, limit int, search string) ([]entity.TravelCost, error) {
	var travelCosts []entity.TravelCost
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ? OR LOWER(code) LIKE ?", 
			"%"+strings.ToLower(search)+"%", 
			"%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Order("id ASC").Find(&travelCosts).Error; err != nil {
		return nil, err
	}
	return travelCosts, nil
}

func (r *travelCostManagementRepo) GetTravelCostByID(id string) (*entity.TravelCost, error) {
	var travelCost entity.TravelCost
	if err := r.db.First(&travelCost, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &travelCost, nil
}

func (r *travelCostManagementRepo) UpdateTravelCost(id string, travelCost *entity.TravelCost) error {
	var existingTravelCost entity.TravelCost
	if err := r.db.First(&existingTravelCost, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingTravelCost).Updates(travelCost).Error
}

func (r *travelCostManagementRepo) DeleteTravelCost(id string) error {
	var travelCost entity.TravelCost
	if err := r.db.First(&travelCost, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&travelCost).Error
}

// Method to get total count of travel costs for pagination
func (r *travelCostManagementRepo) GetTravelCostsCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.TravelCost{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ? OR LOWER(code) LIKE ?", 
			"%"+strings.ToLower(search)+"%", 
			"%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetLastTravelCost retrieves the last travel cost record for code generation
func (r *travelCostManagementRepo) GetLastTravelCost() (*entity.TravelCost, error) {
	var travelCost entity.TravelCost
	if err := r.db.Order("id DESC").First(&travelCost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// If no record is found, return a default travel cost with ID 0
			return &entity.TravelCost{ID: 0}, nil
		}
		return nil, err
	}
	return &travelCost, nil
}
