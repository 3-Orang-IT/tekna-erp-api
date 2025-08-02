package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type ProductManagementUsecase interface {
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
	GetProductsCount(search string) (int64, error) // Method to get total count of products for pagination
}

type productManagementUsecase struct {
	repo adminRepository.ProductManagementRepository
}

func NewProductManagementUsecase(r adminRepository.ProductManagementRepository) ProductManagementUsecase {
	return &productManagementUsecase{repo: r}
}

func (u *productManagementUsecase) CreateProduct(product *entity.Product) error {
	return u.repo.CreateProduct(product)
}

func (u *productManagementUsecase) GetProducts(page, limit int, search string) ([]entity.Product, error) {
	return u.repo.GetProducts(page, limit, search)
}

func (u *productManagementUsecase) GetProductByID(id string) (*entity.Product, error) {
	return u.repo.GetProductByID(id)
}

func (u *productManagementUsecase) UpdateProduct(id string, product *entity.Product) error {
	return u.repo.UpdateProduct(id, product)
}

func (u *productManagementUsecase) DeleteProduct(id string) error {
	return u.repo.DeleteProduct(id)
}

func (u *productManagementUsecase) GetProductCategories() ([]entity.ProductCategory, error) {
	return u.repo.GetProductCategories()
}

func (u *productManagementUsecase) GetSuppliers() ([]entity.Supplier, error) {
	return u.repo.GetSuppliers()
}

func (u *productManagementUsecase) GetBusinessUnits() ([]entity.BusinessUnit, error) {
	return u.repo.GetBusinessUnits()
}

func (u *productManagementUsecase) GetUnits() ([]entity.UnitOfMeasure, error) {
	return u.repo.GetUnits()
}

func (u *productManagementUsecase) GetLastProduct() (*entity.Product, error) {
	return u.repo.GetLastProduct()
}

// GetProductsCount gets the total count of products for pagination
func (u *productManagementUsecase) GetProductsCount(search string) (int64, error) {
	return u.repo.GetProductsCount(search)
}
