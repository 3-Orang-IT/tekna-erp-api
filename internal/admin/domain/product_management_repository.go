package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type ProductManagementRepository interface {
	CreateProduct(product *entity.Product) error
	GetProducts(page, limit int, search string) ([]entity.Product, error)
	GetProductByID(id string) (*entity.Product, error)
	UpdateProduct(id string, product *entity.Product) error
	DeleteProduct(id string) error
	GetProductCategories() ([]entity.ProductCategory, error)
	GetSuppliers() ([]entity.Supplier, error)
	GetBusinessUnits() ([]entity.BusinessUnit, error)
	GetUnits() ([]entity.UnitOfMeasure, error)
	GetLastProduct() (*entity.Product, error)
}
