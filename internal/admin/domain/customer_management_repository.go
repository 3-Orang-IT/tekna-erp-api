package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type CustomerManagementRepository interface {
	CreateCustomer(customer *entity.Customer) error
	GetCustomers(page, limit int, search string) ([]entity.Customer, error)
	GetCustomerByID(id string) (*entity.Customer, error)
	UpdateCustomer(id string, customer *entity.Customer) error
	DeleteCustomer(id string) error
	GetCities(page, limit int, search string) ([]entity.City, error) // For edit page reference
	GetCustomersCount(search string) (int64, error) // Method to get total count of customers for pagination
	GetProvinces(page, limit int, search string) ([]entity.Province, error) // Method for fetching provinces with cities
	GetAreas(page, limit int, search string) ([]entity.Area, error) // Method for fetching areas
	GetUsers(page, limit int, search string) ([]entity.User, error) // Method for fetching users
}
