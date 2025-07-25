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
