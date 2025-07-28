package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type employeeManagementRepo struct {
    db *gorm.DB
}

func NewEmployeeManagementRepository(db *gorm.DB) adminRepository.EmployeeManagementRepository {
    return &employeeManagementRepo{db: db}
}

func (r *employeeManagementRepo) CreateEmployee(employee *entity.Employee) error {
    return r.db.Create(employee).Error
}

func (r *employeeManagementRepo) GetEmployees(page, limit int, search string) ([]entity.Employee, error) {
    var employees []entity.Employee
    offset := (page - 1) * limit
    query := r.db
    if search != "" {
        // Search by n_ip, n_ik, or user.name
        searchStr := "%" + strings.ToLower(search) + "%"
        query = query.Joins("LEFT JOIN users ON users.id = employees.user_id").
            Where("LOWER(employees.nip) LIKE ? OR LOWER(employees.nik) LIKE ? OR LOWER(users.name) LIKE ?", searchStr, searchStr, searchStr)
    }
    if err := query.Preload("User").Preload("JobPosition").Preload("Division").Preload("City").Limit(limit).Offset(offset).Find(&employees).Error; err != nil {
        return nil, err
    }
    return employees, nil
}

func (r *employeeManagementRepo) GetEmployeeByID(id string) (*entity.Employee, error) {
    var employee entity.Employee
    if err := r.db.Preload("User").Preload("JobPosition").Preload("Division").Preload("City").First(&employee, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &employee, nil
}

func (r *employeeManagementRepo) UpdateEmployee(id string, employee *entity.Employee) error {
    var existingEmployee entity.Employee
    if err := r.db.First(&existingEmployee, "id = ?", id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return gorm.ErrRecordNotFound
        }
        return err
    }
    return r.db.Model(&existingEmployee).Updates(employee).Error
}

func (r *employeeManagementRepo) DeleteEmployee(id string) error {
    var employee entity.Employee
    if err := r.db.First(&employee, "id = ?", id).Error; err != nil {
        return err
    }
    return r.db.Delete(&employee).Error
}

// Method to get total count of employees for pagination
func (r *employeeManagementRepo) GetEmployeesCount(search string) (int64, error) {
    var count int64
    query := r.db.Model(&entity.Employee{})
    if search != "" {
        // Search by n_ip, n_ik, or user.name
        searchStr := "%" + strings.ToLower(search) + "%"
        query = query.Joins("LEFT JOIN users ON users.id = employees.user_id").
            Where("LOWER(employees.nip) LIKE ? OR LOWER(employees.nik) LIKE ? OR LOWER(users.name) LIKE ?", searchStr, searchStr, searchStr)
    }
    if err := query.Count(&count).Error; err != nil {
        return 0, err
    }
    return count, nil
}

// GetJobPositions fetches job positions for the edit page
func (r *employeeManagementRepo) GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error) {
    var jobPositions []entity.JobPosition
    offset := (page - 1) * limit
    query := r.db
    if search != "" {
        query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
    }
    if err := query.Limit(limit).Offset(offset).Find(&jobPositions).Error; err != nil {
        return nil, err
    }
    return jobPositions, nil
}

// GetDivisions fetches divisions for the edit page
func (r *employeeManagementRepo) GetDivisions(page, limit int, search string) ([]entity.Division, error) {
    var divisions []entity.Division
    offset := (page - 1) * limit
    query := r.db
    if search != "" {
        query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
    }
    if err := query.Limit(limit).Offset(offset).Find(&divisions).Error; err != nil {
        return nil, err
    }
    return divisions, nil
}

// GetCities fetches cities for the edit page
func (r *employeeManagementRepo) GetCities(page, limit int, search string) ([]entity.City, error) {
    var cities []entity.City
    offset := (page - 1) * limit
    query := r.db
    if search != "" {
        query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
    }
    if err := query.Preload("Province").Limit(limit).Offset(offset).Find(&cities).Error; err != nil {
        return nil, err
    }
    return cities, nil
}

// GetProvinces fetches provinces with their cities for the edit page
func (r *employeeManagementRepo) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
    var provinces []entity.Province
    offset := (page - 1) * limit
    query := r.db
    if search != "" {
        query = query.Where("LOWER(provinces.name) LIKE ?", "%"+strings.ToLower(search)+"%")
    }
    if err := query.Preload("Cities").Limit(limit).Offset(offset).Find(&provinces).Error; err != nil {
        return nil, err
    }
    return provinces, nil
}
