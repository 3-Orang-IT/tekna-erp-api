package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type EmployeeManagementRepository interface {
    CreateEmployee(employee *entity.Employee) error
    GetEmployees(page, limit int, search string) ([]entity.Employee, error)
    GetEmployeeByID(id string) (*entity.Employee, error)
    UpdateEmployee(id string, employee *entity.Employee) error
    DeleteEmployee(id string) error
    // For edit page references, fetch all needed related data if necessary
}
