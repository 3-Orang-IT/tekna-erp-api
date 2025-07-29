package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type supplierManagementRepo struct {
	db *gorm.DB
}

func NewSupplierManagementRepository(db *gorm.DB) adminRepository.SupplierManagementRepository {
	return &supplierManagementRepo{db: db}
}

func (r *supplierManagementRepo) CreateSupplier(supplier *entity.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *supplierManagementRepo) GetSuppliers(page, limit int, search string) ([]entity.Supplier, error) {
	var suppliers []entity.Supplier
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("City.Province").Preload("City").Limit(limit).Offset(offset).Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (r *supplierManagementRepo) GetSupplierByID(id string) (*entity.Supplier, error) {
	var supplier entity.Supplier
	if err := r.db.Preload("City.Province").Preload("City").First(&supplier, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *supplierManagementRepo) UpdateSupplier(id string, supplier *entity.Supplier) error {
	var existingSupplier entity.Supplier
	if err := r.db.First(&existingSupplier, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingSupplier).Updates(supplier).Error
}

func (r *supplierManagementRepo) DeleteSupplier(id string) error {
	var supplier entity.Supplier
	if err := r.db.First(&supplier, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&supplier).Error
}

// Method to get total count of suppliers for pagination
func (r *supplierManagementRepo) GetSuppliersCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Supplier{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetCities fetches cities for the supplier form
func (r *supplierManagementRepo) GetCities(page, limit int, search string) ([]entity.City, error) {
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

// GetProvinces fetches provinces with their cities for the supplier form
func (r *supplierManagementRepo) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
	var provinces []entity.Province
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("Cities").Limit(limit).Offset(offset).Find(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}

// GetUsers fetches users for the supplier form
func (r *supplierManagementRepo) GetUsers(page, limit int, search string) ([]entity.User, error) {
	var users []entity.User
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
