package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type documentManagementRepo struct {
	db *gorm.DB
}

func NewDocumentManagementRepository(db *gorm.DB) adminRepository.DocumentManagementRepository {
	return &documentManagementRepo{db: db}
}

func (r *documentManagementRepo) CreateDocument(document *entity.Document) error {
	return r.db.Create(document).Error
}

func (r *documentManagementRepo) GetDocuments(page, limit int, search string) ([]entity.Document, error) {
	var documents []entity.Document
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("DocumentCategory").Preload("User").Limit(limit).Offset(offset).Order("id ASC").Find(&documents).Error; err != nil {
		return nil, err
	}
	return documents, nil
}

func (r *documentManagementRepo) GetDocumentByID(id string) (*entity.Document, error) {
	var document entity.Document
	if err := r.db.Preload("DocumentCategory").Preload("User").First(&document, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &document, nil
}

func (r *documentManagementRepo) UpdateDocument(id string, document *entity.Document) error {
	var existingDocument entity.Document
	if err := r.db.First(&existingDocument, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingDocument).Updates(document).Error
}

func (r *documentManagementRepo) DeleteDocument(id string) error {
	var document entity.Document
	if err := r.db.First(&document, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&document).Error
}

func (r *documentManagementRepo) GetDocumentCategories() ([]entity.DocumentCategory, error) {
	var documentCategories []entity.DocumentCategory
	if err := r.db.Find(&documentCategories).Error; err != nil {
		return nil, err
	}
	return documentCategories, nil
}

// Method to get total count of documents for pagination
func (r *documentManagementRepo) GetDocumentsCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Document{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
