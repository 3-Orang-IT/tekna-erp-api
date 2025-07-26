package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type businessUnitManagementRepo struct {
	db *gorm.DB
}

func NewBusinessUnitManagementRepository(db *gorm.DB) adminRepository.BusinessUnitManagementRepository {
	return &businessUnitManagementRepo{db: db}
}

func (r *businessUnitManagementRepo) CreateBusinessUnit(businessUnit *entity.BusinessUnit) error {
	return r.db.Create(businessUnit).Error
}

func (r *businessUnitManagementRepo) GetBusinessUnits(page, limit int, search string) ([]entity.BusinessUnit, error) {
	var businessUnits []entity.BusinessUnit
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Find(&businessUnits).Error; err != nil {
		return nil, err
	}
	return businessUnits, nil
}

func (r *businessUnitManagementRepo) GetBusinessUnitByID(id string) (*entity.BusinessUnit, error) {
	var businessUnit entity.BusinessUnit
	if err := r.db.First(&businessUnit, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &businessUnit, nil
}

func (r *businessUnitManagementRepo) UpdateBusinessUnit(id string, businessUnit *entity.BusinessUnit) error {
	var existingBusinessUnit entity.BusinessUnit
	if err := r.db.First(&existingBusinessUnit, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingBusinessUnit).Updates(businessUnit).Error
}

func (r *businessUnitManagementRepo) DeleteBusinessUnit(id string) error {
	var businessUnit entity.BusinessUnit
	if err := r.db.First(&businessUnit, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&businessUnit).Error
}
