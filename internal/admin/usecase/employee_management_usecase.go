package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type EmployeeManagementUsecase interface {
    CreateEmployee(employee *entity.Employee) error
    GetEmployees(page, limit int, search string) ([]entity.Employee, error)
    GetEmployeeByID(id string) (*entity.Employee, error)
    UpdateEmployee(id string, employee *entity.Employee) error
    DeleteEmployee(id string) error
}

type employeeManagementUsecase struct {
    repo adminRepository.EmployeeManagementRepository
}

func NewEmployeeManagementUsecase(r adminRepository.EmployeeManagementRepository) EmployeeManagementUsecase {
    return &employeeManagementUsecase{repo: r}
}

func (u *employeeManagementUsecase) CreateEmployee(employee *entity.Employee) error {
    return u.repo.CreateEmployee(employee)
}

func (u *employeeManagementUsecase) GetEmployees(page, limit int, search string) ([]entity.Employee, error) {
    return u.repo.GetEmployees(page, limit, search)
}

func (u *employeeManagementUsecase) GetEmployeeByID(id string) (*entity.Employee, error) {
    return u.repo.GetEmployeeByID(id)
}

func (u *employeeManagementUsecase) UpdateEmployee(id string, employee *entity.Employee) error {
    return u.repo.UpdateEmployee(id, employee)
}

func (u *employeeManagementUsecase) DeleteEmployee(id string) error {
    return u.repo.DeleteEmployee(id)
}
