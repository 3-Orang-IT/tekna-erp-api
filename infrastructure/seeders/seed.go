package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/domain/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
    // Seed roles
    roles := []entity.Role{
        {Name: "Admin"},
        {Name: "Team Support"},
        {Name: "HRD"},
        {Name: "Supplier"},
    }

    for _, role := range roles {
        var r entity.Role
        result := db.Where("name = ?", role.Name).First(&r)
        if result.Error == gorm.ErrRecordNotFound {
            if err := db.Create(&role).Error; err != nil {
                log.Printf("Gagal membuat role %s: %v", role.Name, err)
            } else {
                log.Printf("Berhasil menambahkan role %s", role.Name)
            }
        }
    }

    // Fetch roles to associate with users
    var adminRole, teamSupportRole entity.Role
    db.Where("name = ?", "Admin").First(&adminRole)
    db.Where("name = ?", "Team Support").First(&teamSupportRole)

    // Seed users
    users := []entity.User{
        {Name: "Admin", Email: "admin@example.com", Password: "hashed_password", RoleID: adminRole.ID},
        {Name: "User", Email: "user@example.com", Password: "hashed_password", RoleID: teamSupportRole.ID},
    }

    for _, user := range users {
        var u entity.User
        result := db.Where("email = ?", user.Email).First(&u)
        if result.Error == gorm.ErrRecordNotFound {
            if err := db.Create(&user).Error; err != nil {
                log.Printf("Gagal membuat user %s: %v", user.Email, err)
            } else {
                log.Printf("Berhasil menambahkan user %s", user.Email)
            }
        }
    }
}