package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type CustomerManagementRepository interface {
	CreateCustomer(customer *entity.Customer) error
	GetCustomers(page, limit int, search string) ([]entity.Customer, error)
	GetCustomerByID(id string) (*entity.Customer, error)
	UpdateCustomer(id string, customer *entity.Customer) error
	DeleteCustomer(id string) error
	GetCities(page, limit int, search string) ([]entity.City, error) // For edit page reference
}
