package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type productCategoryAlternativeManagementRepo struct {
	db *gorm.DB
}

func NewProductCategoryAlternativeManagementRepository(db *gorm.DB) adminRepository.ProductCategoryAlternativeManagementRepository {
	return &productCategoryAlternativeManagementRepo{db: db}
}

func (r *productCategoryAlternativeManagementRepo) CreateProductCategoryAlternative(category *entity.ProductCategoryAlternative) error {
	return r.db.Create(category).Error
}

func (r *productCategoryAlternativeManagementRepo) GetProductCategoryAlternatives(page, limit int, search string) ([]entity.ProductCategoryAlternative, error) {
	var categories []entity.ProductCategoryAlternative
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Order("id ASC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *productCategoryAlternativeManagementRepo) GetProductCategoryAlternativeByID(id string) (*entity.ProductCategoryAlternative, error) {
	var category entity.ProductCategoryAlternative
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *productCategoryAlternativeManagementRepo) UpdateProductCategoryAlternative(id string, category *entity.ProductCategoryAlternative) error {
	var existing entity.ProductCategoryAlternative
	if err := r.db.First(&existing, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existing).Updates(category).Error
}

func (r *productCategoryAlternativeManagementRepo) DeleteProductCategoryAlternative(id string) error {
	var category entity.ProductCategoryAlternative
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&category).Error
}

func (r *productCategoryAlternativeManagementRepo) GetProductCategoryAlternativesCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.ProductCategoryAlternative{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
