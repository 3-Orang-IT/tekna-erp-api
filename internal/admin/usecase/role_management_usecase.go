package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type RoleManagementUsecase interface {
	CreateRole(role *entity.Role) error
	GetRoles(page, limit int, search string) ([]entity.Role, error)
	GetRoleByID(id string) (*entity.Role, error)
	UpdateRole(id string, role *entity.Role) error
	DeleteRole(id string) error
	GetAllMenus(menus *[]entity.Menu) error
	GetRolesCount(search string) (int64, error) // Method to get total count of roles for pagination
}

type roleManagementUsecase struct {
	repo adminRepository.RoleManagementRepository
}

func NewRoleManagementUsecase(r adminRepository.RoleManagementRepository) RoleManagementUsecase {
	return &roleManagementUsecase{repo: r}
}

func (u *roleManagementUsecase) CreateRole(role *entity.Role) error {
	return u.repo.CreateRole(role)
}

func (u *roleManagementUsecase) GetRoles(page, limit int, search string) ([]entity.Role, error) {
	return u.repo.GetRoles(page, limit, search)
}

func (u *roleManagementUsecase) GetRoleByID(id string) (*entity.Role, error) {
	return u.repo.GetRoleByID(id)
}

func (u *roleManagementUsecase) UpdateRole(id string, role *entity.Role) error {
	return u.repo.UpdateRole(id, role)
}

func (u *roleManagementUsecase) DeleteRole(id string) error {
	return u.repo.DeleteRole(id)
}

func (u *roleManagementUsecase) GetAllMenus(menus *[]entity.Menu) error {
	return u.repo.GetAllMenus(menus)
}

// GetRolesCount gets the total count of roles for pagination
func (u *roleManagementUsecase) GetRolesCount(search string) (int64, error) {
	return u.repo.GetRolesCount(search)
}
