package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type productManagementRepo struct {
	db *gorm.DB
}

func NewProductManagementRepository(db *gorm.DB) adminRepository.ProductManagementRepository {
	return &productManagementRepo{db: db}
}

func (r *productManagementRepo) CreateProduct(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *productManagementRepo) GetProducts(page, limit int, search string) ([]entity.Product, error) {
	var products []entity.Product
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("ProductCategory").Preload("Supplier").Preload("BusinessUnit").Preload("Unit").Limit(limit).Offset(offset).Order("id ASC").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productManagementRepo) GetProductByID(id string) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.Preload("ProductCategory").Preload("Supplier").Preload("BusinessUnit").Preload("Unit").First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productManagementRepo) UpdateProduct(id string, product *entity.Product) error {
	var existingProduct entity.Product
	if err := r.db.First(&existingProduct, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingProduct).Updates(product).Error
}

func (r *productManagementRepo) DeleteProduct(id string) error {
	var product entity.Product
	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&product).Error
}

func (r *productManagementRepo) GetProductCategories() ([]entity.ProductCategory, error) {
	var categories []entity.ProductCategory
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *productManagementRepo) GetSuppliers() ([]entity.Supplier, error) {
	var suppliers []entity.Supplier
	if err := r.db.Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (r *productManagementRepo) GetBusinessUnits() ([]entity.BusinessUnit, error) {
	var units []entity.BusinessUnit
	if err := r.db.Find(&units).Error; err != nil {
		return nil, err
	}
	return units, nil
}

func (r *productManagementRepo) GetUnits() ([]entity.UnitOfMeasure, error) {
	var units []entity.UnitOfMeasure
	if err := r.db.Find(&units).Error; err != nil {
		return nil, err
	}
	return units, nil
}

func (r *productManagementRepo) GetLastProduct() (*entity.Product, error) {
	var product entity.Product
	err := r.db.Order("id desc").First(&product).Error
	if err == gorm.ErrRecordNotFound {
		return &entity.Product{ID: 0}, nil
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Method to get total count of products for pagination
func (r *productManagementRepo) GetProductsCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Product{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
