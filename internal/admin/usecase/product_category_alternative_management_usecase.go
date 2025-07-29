package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type ProductCategoryAlternativeManagementUsecase interface {
	CreateProductCategoryAlternative(category *entity.ProductCategoryAlternative) error
	GetProductCategoryAlternatives(page, limit int, search string) ([]entity.ProductCategoryAlternative, error)
	GetProductCategoryAlternativeByID(id string) (*entity.ProductCategoryAlternative, error)
	UpdateProductCategoryAlternative(id string, category *entity.ProductCategoryAlternative) error
	DeleteProductCategoryAlternative(id string) error
	GetProductCategoryAlternativesCount(search string) (int64, error)
}

type productCategoryAlternativeManagementUsecase struct {
	repo adminRepository.ProductCategoryAlternativeManagementRepository
}

func NewProductCategoryAlternativeManagementUsecase(r adminRepository.ProductCategoryAlternativeManagementRepository) ProductCategoryAlternativeManagementUsecase {
	return &productCategoryAlternativeManagementUsecase{repo: r}
}

func (u *productCategoryAlternativeManagementUsecase) CreateProductCategoryAlternative(category *entity.ProductCategoryAlternative) error {
	return u.repo.CreateProductCategoryAlternative(category)
}

func (u *productCategoryAlternativeManagementUsecase) GetProductCategoryAlternatives(page, limit int, search string) ([]entity.ProductCategoryAlternative, error) {
	return u.repo.GetProductCategoryAlternatives(page, limit, search)
}

func (u *productCategoryAlternativeManagementUsecase) GetProductCategoryAlternativeByID(id string) (*entity.ProductCategoryAlternative, error) {
	return u.repo.GetProductCategoryAlternativeByID(id)
}

func (u *productCategoryAlternativeManagementUsecase) UpdateProductCategoryAlternative(id string, category *entity.ProductCategoryAlternative) error {
	return u.repo.UpdateProductCategoryAlternative(id, category)
}

func (u *productCategoryAlternativeManagementUsecase) DeleteProductCategoryAlternative(id string) error {
	return u.repo.DeleteProductCategoryAlternative(id)
}

func (u *productCategoryAlternativeManagementUsecase) GetProductCategoryAlternativesCount(search string) (int64, error) {
	return u.repo.GetProductCategoryAlternativesCount(search)
}
