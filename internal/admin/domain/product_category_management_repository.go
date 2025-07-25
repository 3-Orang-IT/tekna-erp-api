package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type ProductCategoryManagementRepository interface {
	CreateProductCategory(category *entity.ProductCategory) error
	GetProductCategories(page, limit int, search string) ([]entity.ProductCategory, error)
	GetProductCategoryByID(id string) (*entity.ProductCategory, error)
	UpdateProductCategory(id string, category *entity.ProductCategory) error
	DeleteProductCategory(id string) error
}
