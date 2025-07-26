package adminRepositoryImpl

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userManagementRepo struct {
	db *gorm.DB
}

func NewUserManagementRepository(db *gorm.DB) adminRepository.UserManagementRepository {
	return &userManagementRepo{db: db}
}

func (r *userManagementRepo) CreateUser(user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return r.db.Create(user).Error
}

func (r *userManagementRepo) GetUsers(page, limit int, search string) ([]entity.User, error) {
	var users []entity.User
	offset := (page - 1) * limit
	query := r.db.Preload("Role").Limit(limit).Offset(offset)
	if search != "" {
		query = query.Where("username LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userManagementRepo) GetUserByID(id string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Preload("Role").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userManagementRepo) UpdateUser(id string, user *entity.User) error {
	var existingUser entity.User
	if err := r.db.First(&existingUser, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound // Return specific error if user ID is not found
		}
		return err
	}

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return r.db.Model(&existingUser).Updates(user).Error
}

func (r *userManagementRepo) DeleteUser(id string) error {
	var user entity.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return err // Return error if user not found
	}

	return r.db.Delete(&user).Error // Proceed to delete if user exists
}

func (r *userManagementRepo) GetAllRoles() ([]entity.Role, error) {
	var roles []entity.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}