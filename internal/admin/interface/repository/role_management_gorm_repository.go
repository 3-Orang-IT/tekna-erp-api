package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type roleManagementRepo struct {
	db *gorm.DB
}

func NewRoleManagementRepository(db *gorm.DB) adminRepository.RoleManagementRepository {
	return &roleManagementRepo{db: db}
}

func (r *roleManagementRepo) CreateRole(role *entity.Role) error {
	return r.db.Create(role).Error
}

func (r *roleManagementRepo) GetRoles(page, limit int, search string) ([]entity.Role, error) {
	var roles []entity.Role
	offset := (page - 1) * limit
	query := r.db.Preload("Menus").Limit(limit).Offset(offset)

	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	if err := query.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleManagementRepo) GetRoleByID(id string) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.Preload("Menus").First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleManagementRepo) UpdateRole(id string, role *entity.Role) error {
	var existingRole entity.Role
	if err := r.db.First(&existingRole, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingRole).Updates(role).Error
}

func (r *roleManagementRepo) DeleteRole(id string) error {
	var role entity.Role
	if err := r.db.First(&role, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&role).Error
}

func (r *roleManagementRepo) GetAllMenus(menus *[]entity.Menu) error {
	if err := r.db.Find(menus).Error; err != nil {
		return err
	}
	return nil
}
