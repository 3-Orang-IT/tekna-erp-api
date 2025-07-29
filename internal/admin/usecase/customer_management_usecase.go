package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type CustomerManagementUsecase interface {
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

type customerManagementUsecase struct {
	repo adminRepository.CustomerManagementRepository
}

func NewCustomerManagementUsecase(r adminRepository.CustomerManagementRepository) CustomerManagementUsecase {
	return &customerManagementUsecase{repo: r}
}

func (u *customerManagementUsecase) CreateCustomer(customer *entity.Customer) error {
	return u.repo.CreateCustomer(customer)
}

func (u *customerManagementUsecase) GetCustomers(page, limit int, search string) ([]entity.Customer, error) {
	return u.repo.GetCustomers(page, limit, search)
}

func (u *customerManagementUsecase) GetCustomerByID(id string) (*entity.Customer, error) {
	return u.repo.GetCustomerByID(id)
}

func (u *customerManagementUsecase) UpdateCustomer(id string, customer *entity.Customer) error {
	return u.repo.UpdateCustomer(id, customer)
}

func (u *customerManagementUsecase) DeleteCustomer(id string) error {
	return u.repo.DeleteCustomer(id)
}

func (u *customerManagementUsecase) GetCities(page, limit int, search string) ([]entity.City, error) {
	return u.repo.GetCities(page, limit, search)
}

// GetCustomersCount gets the total count of customers for pagination
func (u *customerManagementUsecase) GetCustomersCount(search string) (int64, error) {
	return u.repo.GetCustomersCount(search)
}

// GetProvinces fetches provinces with their cities for the customer form
func (u *customerManagementUsecase) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
	return u.repo.GetProvinces(page, limit, search)
}

// GetAreas fetches areas for the customer form
func (u *customerManagementUsecase) GetAreas(page, limit int, search string) ([]entity.Area, error) {
	return u.repo.GetAreas(page, limit, search)
}

// GetUsers fetches users for the customer form
func (u *customerManagementUsecase) GetUsers(page, limit int, search string) ([]entity.User, error) {
	return u.repo.GetUsers(page, limit, search)
}
