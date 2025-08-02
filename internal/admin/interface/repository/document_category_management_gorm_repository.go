package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type documentCategoryManagementRepo struct {
	db *gorm.DB
}

func NewDocumentCategoryManagementRepository(db *gorm.DB) adminRepository.DocumentCategoryManagementRepository {
	return &documentCategoryManagementRepo{db: db}
}

func (r *documentCategoryManagementRepo) CreateDocumentCategory(documentCategory *entity.DocumentCategory) error {
	return r.db.Create(documentCategory).Error
}

func (r *documentCategoryManagementRepo) GetDocumentCategories(page, limit int, search string) ([]entity.DocumentCategory, error) {
	var documentCategories []entity.DocumentCategory
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Order("id ASC").Find(&documentCategories).Error; err != nil {
		return nil, err
	}
	return documentCategories, nil
}

func (r *documentCategoryManagementRepo) GetDocumentCategoryByID(id string) (*entity.DocumentCategory, error) {
	var documentCategory entity.DocumentCategory
	if err := r.db.First(&documentCategory, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &documentCategory, nil
}

func (r *documentCategoryManagementRepo) UpdateDocumentCategory(id string, documentCategory *entity.DocumentCategory) error {
	var existingDocumentCategory entity.DocumentCategory
	if err := r.db.First(&existingDocumentCategory, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingDocumentCategory).Updates(documentCategory).Error
}

func (r *documentCategoryManagementRepo) DeleteDocumentCategory(id string) error {
	var documentCategory entity.DocumentCategory
	if err := r.db.First(&documentCategory, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&documentCategory).Error
}

// Method to get total count of document categories for pagination
func (r *documentCategoryManagementRepo) GetDocumentCategoriesCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.DocumentCategory{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
