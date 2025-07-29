package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type DocumentManagementUsecase interface {
	CreateDocument(document *entity.Document) error
	GetDocuments(page, limit int, search string) ([]entity.Document, error)
	GetDocumentByID(id string) (*entity.Document, error)
	UpdateDocument(id string, document *entity.Document) error
	DeleteDocument(id string) error
	GetDocumentCategories() ([]entity.DocumentCategory, error)
	GetDocumentsCount(search string) (int64, error) // Method to get total count of documents for pagination
}

type documentManagementUsecase struct {
	repo adminRepository.DocumentManagementRepository
}

func NewDocumentManagementUsecase(r adminRepository.DocumentManagementRepository) DocumentManagementUsecase {
	return &documentManagementUsecase{repo: r}
}

func (u *documentManagementUsecase) CreateDocument(document *entity.Document) error {
	return u.repo.CreateDocument(document)
}

func (u *documentManagementUsecase) GetDocuments(page, limit int, search string) ([]entity.Document, error) {
	return u.repo.GetDocuments(page, limit, search)
}

func (u *documentManagementUsecase) GetDocumentByID(id string) (*entity.Document, error) {
	return u.repo.GetDocumentByID(id)
}

func (u *documentManagementUsecase) UpdateDocument(id string, document *entity.Document) error {
	return u.repo.UpdateDocument(id, document)
}

func (u *documentManagementUsecase) DeleteDocument(id string) error {
	return u.repo.DeleteDocument(id)
}

func (u *documentManagementUsecase) GetDocumentCategories() ([]entity.DocumentCategory, error) {
	return u.repo.GetDocumentCategories()
}

// GetDocumentsCount gets the total count of documents for pagination
func (u *documentManagementUsecase) GetDocumentsCount(search string) (int64, error) {
	return u.repo.GetDocumentsCount(search)
}
