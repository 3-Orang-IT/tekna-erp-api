package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type unitOfMeasureManagementRepo struct {
	db *gorm.DB
}

func NewUnitOfMeasureManagementRepository(db *gorm.DB) adminRepository.UnitOfMeasureManagementRepository {
	return &unitOfMeasureManagementRepo{db: db}
}

func (r *unitOfMeasureManagementRepo) CreateUnitOfMeasure(unitOfMeasure *entity.UnitOfMeasure) error {
	return r.db.Create(unitOfMeasure).Error
}

func (r *unitOfMeasureManagementRepo) GetUnitOfMeasures(page, limit int, search string) ([]entity.UnitOfMeasure, error) {
	var unitOfMeasures []entity.UnitOfMeasure
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Find(&unitOfMeasures).Error; err != nil {
		return nil, err
	}
	return unitOfMeasures, nil
}

func (r *unitOfMeasureManagementRepo) GetUnitOfMeasureByID(id string) (*entity.UnitOfMeasure, error) {
	var unitOfMeasure entity.UnitOfMeasure
	if err := r.db.First(&unitOfMeasure, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &unitOfMeasure, nil
}

func (r *unitOfMeasureManagementRepo) UpdateUnitOfMeasure(id string, unitOfMeasure *entity.UnitOfMeasure) error {
	var existingUnitOfMeasure entity.UnitOfMeasure
	if err := r.db.First(&existingUnitOfMeasure, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingUnitOfMeasure).Updates(unitOfMeasure).Error
}

func (r *unitOfMeasureManagementRepo) DeleteUnitOfMeasure(id string) error {
	var unitOfMeasure entity.UnitOfMeasure
	if err := r.db.First(&unitOfMeasure, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&unitOfMeasure).Error
}
