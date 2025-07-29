package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type EmployeeManagementUsecase interface {
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

// GetEmployeesCount gets the total count of employees for pagination
func (u *employeeManagementUsecase) GetEmployeesCount(search string) (int64, error) {
    return u.repo.GetEmployeesCount(search)
}

// GetJobPositions fetches job positions for the edit page
func (u *employeeManagementUsecase) GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error) {
    return u.repo.GetJobPositions(page, limit, search)
}

// GetDivisions fetches divisions for the edit page
func (u *employeeManagementUsecase) GetDivisions(page, limit int, search string) ([]entity.Division, error) {
    return u.repo.GetDivisions(page, limit, search)
}

// GetCities fetches cities for the edit page
func (u *employeeManagementUsecase) GetCities(page, limit int, search string) ([]entity.City, error) {
    return u.repo.GetCities(page, limit, search)
}

// GetProvinces fetches provinces with their cities for the edit page
func (u *employeeManagementUsecase) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
    return u.repo.GetProvinces(page, limit, search)
}

// CreateEmployeeWithUser creates a new user and assigns them as an employee
func (u *employeeManagementUsecase) CreateEmployeeWithUser(user *entity.User, employee *entity.Employee, roleIDs []uint) error {
    return u.repo.CreateEmployeeWithUser(user, employee, roleIDs)
}

// GetUsers fetches users that can be assigned to employees
func (u *employeeManagementUsecase) GetUsers(page, limit int, search string) ([]entity.User, error) {
    return u.repo.GetUsers(page, limit, search)
}

// GetAreas fetches areas for employee assignment
func (u *employeeManagementUsecase) GetAreas(page, limit int, search string) ([]entity.Area, error) {
    return u.repo.GetAreas(page, limit, search)
}
