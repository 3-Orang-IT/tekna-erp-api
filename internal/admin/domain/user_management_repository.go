package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type UserManagementRepository interface {
	CreateUser(user *entity.User) error
	GetUsers(page, limit int, search string) ([]entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	UpdateUser(id string, user *entity.User) error
	DeleteUser(id string) error
	GetAllRoles() ([]entity.Role, error)
	// Method to get total count of users for pagination
	GetUsersCount(search string) (int64, error)
}



