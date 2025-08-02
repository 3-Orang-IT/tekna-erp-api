package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type EmployeeManagementRepository interface {
    CreateEmployee(employee *entity.Employee) error
    CreateEmployeeWithUser(user *entity.User, employee *entity.Employee, roleIDs []uint) error
    GetEmployees(page, limit int, search string) ([]entity.Employee, error)
    GetEmployeeByID(id string) (*entity.Employee, error)
    UpdateEmployee(id string, employee *entity.Employee) error
    DeleteEmployee(id string) error
    // Method to get total count of employees for pagination
    GetEmployeesCount(search string) (int64, error)
    // Methods for fetching reference data for edit page
    GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error)
    GetDivisions(page, limit int, search string) ([]entity.Division, error)
    GetCities(page, limit int, search string) ([]entity.City, error)
    GetProvinces(page, limit int, search string) ([]entity.Province, error)
    // Method to get areas for employee assignment
    GetAreas(page, limit int, search string) ([]entity.Area, error)
    // Method to get users that can be assigned to employees
    GetUsers(page, limit int, search string) ([]entity.User, error)
}
