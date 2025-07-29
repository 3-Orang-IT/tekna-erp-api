package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type ProductCategoryAlternativeManagementRepository interface {
	CreateProductCategoryAlternative(category *entity.ProductCategoryAlternative) error
	GetProductCategoryAlternatives(page, limit int, search string) ([]entity.ProductCategoryAlternative, error)
	GetProductCategoryAlternativeByID(id string) (*entity.ProductCategoryAlternative, error)
	UpdateProductCategoryAlternative(id string, category *entity.ProductCategoryAlternative) error
	DeleteProductCategoryAlternative(id string) error
	GetProductCategoryAlternativesCount(search string) (int64, error)
}
