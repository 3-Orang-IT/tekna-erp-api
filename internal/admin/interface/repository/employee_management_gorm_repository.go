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
