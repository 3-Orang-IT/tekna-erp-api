package repository

import (
	"fmt"

	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/domain/entity"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/domain/repository"
	"gorm.io/gorm"
)

type authRepo struct {
    db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
    return &authRepo{db: db}
}

func (r *authRepo) Register(user *entity.User) error {
    var adminRole entity.Role
    if err := r.db.Where("name = ?", "Admin").First(&adminRole).Error; err != nil {
        return fmt.Errorf("default role not found: %w", err)
    }

    // Set user role
    user.Role = []entity.Role{adminRole}
    return r.db.Create(user).Error
}

func (r *authRepo) FindByEmail(email string) (*entity.User, error) {
    var user entity.User
    if err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *authRepo) GetMenusByRoleID(roleID uint) ([]entity.Menu, error) {
    var role entity.Role
    if err := r.db.Preload("Menus").First(&role, roleID).Error; err != nil {
        return nil, err
    }
    return role.Menus, nil
}
