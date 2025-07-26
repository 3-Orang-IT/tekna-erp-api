package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type UserManagementUsecase interface {
	CreateUser(user *entity.User) error
	GetUsers(page, limit int, search string) ([]entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	UpdateUser(id string, user *entity.User) error
	DeleteUser(id string) error
	GetAllRoles() ([]entity.Role, error)
}

type userManagementUsecase struct {
	repo adminRepository.UserManagementRepository
}

func NewUserManagementUsecase(r adminRepository.UserManagementRepository) UserManagementUsecase {
	return &userManagementUsecase{repo: r}
}

func (u *userManagementUsecase) CreateUser(user *entity.User) error {
	return u.repo.CreateUser(user)
}

func (u *userManagementUsecase) GetUsers(page, limit int, search string) ([]entity.User, error) {
	return u.repo.GetUsers(page, limit, search)
}

func (u *userManagementUsecase) GetUserByID(id string) (*entity.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *userManagementUsecase) UpdateUser(id string, user *entity.User) error {
	return u.repo.UpdateUser(id, user)
}

func (u *userManagementUsecase) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}

func (u *userManagementUsecase) GetAllRoles() ([]entity.Role, error) {
	return u.repo.GetAllRoles()
}