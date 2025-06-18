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

func (r *authRepo) FindByUsername(username string) (*entity.User, error) {
    var user entity.User
    if err := r.db.Preload("Role").Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *authRepo) GetMenusByUserID(userID uint) ([]entity.Menu, error) {
    var menus []entity.Menu

    err := r.db.
        Distinct().
        Table("menus").
        Joins("JOIN role_menus ON role_menus.menu_id = menus.id").
        Joins("JOIN user_roles ON user_roles.role_id = role_menus.role_id").
        Where("user_roles.user_id = ?", userID).
        Find(&menus).Error

    if err != nil {
        return nil, err
    }

    return menus, nil
}


