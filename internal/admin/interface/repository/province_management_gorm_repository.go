package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type provinceManagementRepo struct {
	db *gorm.DB
}

func NewProvinceManagementRepository(db *gorm.DB) adminRepository.ProvinceManagementRepository {
	return &provinceManagementRepo{db: db}
}

func (r *provinceManagementRepo) CreateProvince(province *entity.Province) error {
	return r.db.Create(province).Error
}

func (r *provinceManagementRepo) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
	var provinces []entity.Province
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Order("id ASC").Find(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r *provinceManagementRepo) GetProvincesCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Province{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *provinceManagementRepo) GetProvinceByID(id string) (*entity.Province, error) {
	var province entity.Province
	if err := r.db.First(&province, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &province, nil
}

func (r *provinceManagementRepo) UpdateProvince(id string, province *entity.Province) error {
	var existingProvince entity.Province
	if err := r.db.First(&existingProvince, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingProvince).Updates(province).Error
}

func (r *provinceManagementRepo) DeleteProvince(id string) error {
	var province entity.Province
	if err := r.db.First(&province, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&province).Error
}
