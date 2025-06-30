package adminRepository

import (
	repository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type userManagementRepo struct {
	db *gorm.DB
}

func NewUserManagementRepository(db *gorm.DB) repository.UserManagementRepository {
	return &userManagementRepo{db: db}
}

func (r *userManagementRepo) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userManagementRepo) GetUsers() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userManagementRepo) GetUserByID(id string) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userManagementRepo) UpdateUser(id string, user *entity.User) error {
	return r.db.Model(&entity.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userManagementRepo) DeleteUser(id string) error {
	var user entity.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return err // Return error if user not found
	}
	return r.db.Delete(&user).Error // Proceed to delete if user exists
}