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
