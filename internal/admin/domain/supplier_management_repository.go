package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type SupplierManagementRepository interface {
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
