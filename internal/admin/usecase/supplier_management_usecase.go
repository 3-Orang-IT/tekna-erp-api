package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type SupplierManagementUsecase interface {
	CreateSupplier(supplier *entity.Supplier) error
	GetSuppliers(page, limit int, search string) ([]entity.Supplier, error)
	GetSupplierByID(id string) (*entity.Supplier, error)
	UpdateSupplier(id string, supplier *entity.Supplier) error
	DeleteSupplier(id string) error
	GetSuppliersCount(search string) (int64, error) // Method to get total count of suppliers for pagination
	GetCities(page, limit int, search string) ([]entity.City, error) // Method for fetching cities
	GetProvinces(page, limit int, search string) ([]entity.Province, error) // Method for fetching provinces with cities
	GetUsers(page, limit int, search string) ([]entity.User, error) // Method for fetching users
}

type supplierManagementUsecase struct {
	repo adminRepository.SupplierManagementRepository
}

func NewSupplierManagementUsecase(r adminRepository.SupplierManagementRepository) SupplierManagementUsecase {
	return &supplierManagementUsecase{repo: r}
}

func (u *supplierManagementUsecase) CreateSupplier(supplier *entity.Supplier) error {
	return u.repo.CreateSupplier(supplier)
}

func (u *supplierManagementUsecase) GetSuppliers(page, limit int, search string) ([]entity.Supplier, error) {
	return u.repo.GetSuppliers(page, limit, search)
}

func (u *supplierManagementUsecase) GetSupplierByID(id string) (*entity.Supplier, error) {
	return u.repo.GetSupplierByID(id)
}

func (u *supplierManagementUsecase) UpdateSupplier(id string, supplier *entity.Supplier) error {
	return u.repo.UpdateSupplier(id, supplier)
}

func (u *supplierManagementUsecase) DeleteSupplier(id string) error {
	return u.repo.DeleteSupplier(id)
}

// GetSuppliersCount gets the total count of suppliers for pagination
func (u *supplierManagementUsecase) GetSuppliersCount(search string) (int64, error) {
	return u.repo.GetSuppliersCount(search)
}

// GetCities gets the cities for the supplier form
func (u *supplierManagementUsecase) GetCities(page, limit int, search string) ([]entity.City, error) {
	return u.repo.GetCities(page, limit, search)
}

// GetProvinces gets the provinces with their cities for the supplier form
func (u *supplierManagementUsecase) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
	return u.repo.GetProvinces(page, limit, search)
}

// GetUsers gets the users for the supplier form
func (u *supplierManagementUsecase) GetUsers(page, limit int, search string) ([]entity.User, error) {
	return u.repo.GetUsers(page, limit, search)
}
