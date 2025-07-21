package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type UserManagementRepository interface {
	CreateUser(user *entity.User) error
	GetUsers(page, limit int) ([]entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	UpdateUser(id string, user *entity.User) error
	DeleteUser(id string) error
}



