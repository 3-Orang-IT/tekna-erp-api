package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type RoleManagementUsecase interface {
	CreateRole(role *entity.Role) error
	GetRoles(page, limit int) ([]entity.Role, error)
	GetRoleByID(id string) (*entity.Role, error)
	UpdateRole(id string, role *entity.Role) error
	DeleteRole(id string) error
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

func (u *roleManagementUsecase) GetRoles(page, limit int) ([]entity.Role, error) {
	return u.repo.GetRoles(page, limit)
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
