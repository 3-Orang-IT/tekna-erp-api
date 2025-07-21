package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type RoleManagementRepository interface {
	CreateRole(role *entity.Role) error
	GetRoles(page, limit int) ([]entity.Role, error)
	GetRoleByID(id string) (*entity.Role, error)
	UpdateRole(id string, role *entity.Role) error
	DeleteRole(id string) error
}
