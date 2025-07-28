package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type customerManagementRepo struct {
	db *gorm.DB
}

func NewCustomerManagementRepository(db *gorm.DB) adminRepository.CustomerManagementRepository {
	return &customerManagementRepo{db: db}
}

func (r *customerManagementRepo) CreateCustomer(customer *entity.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerManagementRepo) GetCustomers(page, limit int, search string) ([]entity.Customer, error) {
	var customers []entity.Customer
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(invoice_name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.
		Preload("City.Province").
		Preload("Area").
		Limit(limit).
		Offset(offset).
		Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *customerManagementRepo) GetCustomerByID(id string) (*entity.Customer, error) {
	var customer entity.Customer
	if err := r.db.
		Preload("City.Province").
		Preload("Area").
		Preload("User").
		First(&customer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerManagementRepo) UpdateCustomer(id string, customer *entity.Customer) error {
	var existingCustomer entity.Customer
	if err := r.db.First(&existingCustomer, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingCustomer).Updates(customer).Error
}

func (r *customerManagementRepo) DeleteCustomer(id string) error {
	var customer entity.Customer
	if err := r.db.First(&customer, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&customer).Error
}

func (r *customerManagementRepo) GetCities(page, limit int, search string) ([]entity.City, error) {
	var cities []entity.City
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}
