package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type productCategoryManagementRepo struct {
	db *gorm.DB
}

func NewProductCategoryManagementRepository(db *gorm.DB) adminRepository.ProductCategoryManagementRepository {
	return &productCategoryManagementRepo{db: db}
}

func (r *productCategoryManagementRepo) CreateProductCategory(category *entity.ProductCategory) error {
	return r.db.Create(category).Error
}

func (r *productCategoryManagementRepo) GetProductCategories(page, limit int, search string) ([]entity.ProductCategory, error) {
	var categories []entity.ProductCategory
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *productCategoryManagementRepo) GetProductCategoryByID(id string) (*entity.ProductCategory, error) {
	var category entity.ProductCategory
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *productCategoryManagementRepo) UpdateProductCategory(id string, category *entity.ProductCategory) error {
	var existing entity.ProductCategory
	if err := r.db.First(&existing, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existing).Updates(category).Error
}

func (r *productCategoryManagementRepo) DeleteProductCategory(id string) error {
	var category entity.ProductCategory
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&category).Error
}

// Method to get total count of product categories for pagination
func (r *productCategoryManagementRepo) GetProductCategoriesCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.ProductCategory{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
