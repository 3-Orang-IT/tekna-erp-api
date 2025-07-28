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
