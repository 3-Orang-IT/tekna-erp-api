package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type DocumentCategoryManagementUsecase interface {
	CreateDocumentCategory(documentCategory *entity.DocumentCategory) error
	GetDocumentCategories(page, limit int, search string) ([]entity.DocumentCategory, error)
	GetDocumentCategoryByID(id string) (*entity.DocumentCategory, error)
	UpdateDocumentCategory(id string, documentCategory *entity.DocumentCategory) error
	DeleteDocumentCategory(id string) error
	GetDocumentCategoriesCount(search string) (int64, error) // Method to get total count of document categories for pagination
}

type documentCategoryManagementUsecase struct {
	repo adminRepository.DocumentCategoryManagementRepository
}

func NewDocumentCategoryManagementUsecase(r adminRepository.DocumentCategoryManagementRepository) DocumentCategoryManagementUsecase {
	return &documentCategoryManagementUsecase{repo: r}
}

func (u *documentCategoryManagementUsecase) CreateDocumentCategory(documentCategory *entity.DocumentCategory) error {
	return u.repo.CreateDocumentCategory(documentCategory)
}

func (u *documentCategoryManagementUsecase) GetDocumentCategories(page, limit int, search string) ([]entity.DocumentCategory, error) {
	return u.repo.GetDocumentCategories(page, limit, search)
}

func (u *documentCategoryManagementUsecase) GetDocumentCategoryByID(id string) (*entity.DocumentCategory, error) {
	return u.repo.GetDocumentCategoryByID(id)
}

func (u *documentCategoryManagementUsecase) UpdateDocumentCategory(id string, documentCategory *entity.DocumentCategory) error {
	return u.repo.UpdateDocumentCategory(id, documentCategory)
}

func (u *documentCategoryManagementUsecase) DeleteDocumentCategory(id string) error {
	return u.repo.DeleteDocumentCategory(id)
}

// GetDocumentCategoriesCount gets the total count of document categories for pagination
func (u *documentCategoryManagementUsecase) GetDocumentCategoriesCount(search string) (int64, error) {
	return u.repo.GetDocumentCategoriesCount(search)
}
