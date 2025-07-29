package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type DocumentCategoryManagementRepository interface {
	CreateDocumentCategory(documentCategory *entity.DocumentCategory) error
	GetDocumentCategories(page, limit int, search string) ([]entity.DocumentCategory, error)
	GetDocumentCategoryByID(id string) (*entity.DocumentCategory, error)
	UpdateDocumentCategory(id string, documentCategory *entity.DocumentCategory) error
	DeleteDocumentCategory(id string) error
	// Method to get total count of document categories for pagination
	GetDocumentCategoriesCount(search string) (int64, error)
}
