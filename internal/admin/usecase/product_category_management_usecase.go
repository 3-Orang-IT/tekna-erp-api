package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type ProductCategoryManagementUsecase interface {
	CreateProductCategory(category *entity.ProductCategory) error
	GetProductCategories(page, limit int, search string) ([]entity.ProductCategory, error)
	GetProductCategoryByID(id string) (*entity.ProductCategory, error)
	UpdateProductCategory(id string, category *entity.ProductCategory) error
	DeleteProductCategory(id string) error
	GetProductCategoriesCount(search string) (int64, error) // Method to get total count of product categories for pagination
}

type productCategoryManagementUsecase struct {
	repo adminRepository.ProductCategoryManagementRepository
}

func NewProductCategoryManagementUsecase(r adminRepository.ProductCategoryManagementRepository) ProductCategoryManagementUsecase {
	return &productCategoryManagementUsecase{repo: r}
}

func (u *productCategoryManagementUsecase) CreateProductCategory(category *entity.ProductCategory) error {
	return u.repo.CreateProductCategory(category)
}

func (u *productCategoryManagementUsecase) GetProductCategories(page, limit int, search string) ([]entity.ProductCategory, error) {
	return u.repo.GetProductCategories(page, limit, search)
}

func (u *productCategoryManagementUsecase) GetProductCategoryByID(id string) (*entity.ProductCategory, error) {
	return u.repo.GetProductCategoryByID(id)
}

func (u *productCategoryManagementUsecase) UpdateProductCategory(id string, category *entity.ProductCategory) error {
	return u.repo.UpdateProductCategory(id, category)
}

func (u *productCategoryManagementUsecase) DeleteProductCategory(id string) error {
	return u.repo.DeleteProductCategory(id)
}

// GetProductCategoriesCount gets the total count of product categories for pagination
func (u *productCategoryManagementUsecase) GetProductCategoriesCount(search string) (int64, error) {
	return u.repo.GetProductCategoriesCount(search)
}
