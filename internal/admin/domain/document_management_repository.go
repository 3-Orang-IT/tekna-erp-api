package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type DocumentManagementRepository interface {
	CreateDocument(document *entity.Document) error
	GetDocuments(page, limit int, search string) ([]entity.Document, error)
	GetDocumentByID(id string) (*entity.Document, error)
	UpdateDocument(id string, document *entity.Document) error
	DeleteDocument(id string) error
	GetDocumentCategories() ([]entity.DocumentCategory, error)
	// Method to get total count of documents for pagination
	GetDocumentsCount(search string) (int64, error)
}
