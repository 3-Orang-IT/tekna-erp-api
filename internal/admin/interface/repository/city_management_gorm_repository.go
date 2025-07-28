package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type cityManagementRepo struct {
	db *gorm.DB
}

func NewCityManagementRepository(db *gorm.DB) adminRepository.CityManagementRepository {
	return &cityManagementRepo{db: db}
}

func (r *cityManagementRepo) CreateCity(city *entity.City) error {
	return r.db.Create(city).Error
}

func (r *cityManagementRepo) GetCities(page, limit int, search string) ([]entity.City, error) {
	var cities []entity.City
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("Province").Limit(limit).Offset(offset).Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *cityManagementRepo) GetCitiesCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.City{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *cityManagementRepo) GetCityByID(id string) (*entity.City, error) {
	var city entity.City
	if err := r.db.Preload("Province").First(&city, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *cityManagementRepo) UpdateCity(id string, city *entity.City) error {
	var existingCity entity.City
	if err := r.db.First(&existingCity, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingCity).Updates(city).Error
}

func (r *cityManagementRepo) DeleteCity(id string) error {
	var city entity.City
	if err := r.db.First(&city, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&city).Error
}

func (r *cityManagementRepo) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
	var provinces []entity.Province
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Find(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}
