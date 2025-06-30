package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
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

	if err := db.Find(&roles).Error; err != nil {
		log.Printf("Gagal mengambil roles: %v", err)
		return
	}

	user := entity.User{
		Username: "admin_user",
		Password: "securepassword", // Replace with hashed password in production
		Name:     "Admin User",
		Email:    "admin@example.com",
		Telp:     "123456789",
		Status:   "active",
		Role:     roles,
	}

	var existingUser entity.User
	result := db.Where("username = ?", user.Username).First(&existingUser)
	if result.Error == gorm.ErrRecordNotFound {
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Gagal membuat user %s: %v", user.Username, err)
		} else {
			log.Printf("Berhasil menambahkan user %s dengan semua role", user.Username)
		}
	}
}