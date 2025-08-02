package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type divisionManagementRepo struct {
	db *gorm.DB
}

func NewDivisionManagementRepository(db *gorm.DB) adminRepository.DivisionManagementRepository {
	return &divisionManagementRepo{db: db}
}

func (r *divisionManagementRepo) CreateDivision(division *entity.Division) error {
	return r.db.Create(division).Error
}

func (r *divisionManagementRepo) GetDivisions(page, limit int, search string) ([]entity.Division, error) {
	var divisions []entity.Division
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Order("id ASC").Find(&divisions).Error; err != nil {
		return nil, err
	}
	return divisions, nil
}

func (r *divisionManagementRepo) GetDivisionByID(id string) (*entity.Division, error) {
	var division entity.Division
	if err := r.db.First(&division, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &division, nil
}

func (r *divisionManagementRepo) UpdateDivision(id string, division *entity.Division) error {
	var existingDivision entity.Division
	if err := r.db.First(&existingDivision, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingDivision).Updates(division).Error
}

func (r *divisionManagementRepo) DeleteDivision(id string) error {
	var division entity.Division
	if err := r.db.First(&division, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&division).Error
}

// Method to get total count of divisions for pagination
func (r *divisionManagementRepo) GetDivisionsCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Division{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
