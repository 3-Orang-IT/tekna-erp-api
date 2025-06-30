package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Seed roles
	rolesList := []entity.Role{
		{Name: "Admin"},
		{Name: "Team Support"},
		{Name: "HRD"},
		{Name: "Supplier"},
	}

	for _, role := range rolesList {
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

	// Seed user with all roles
	var roles []entity.Role
	if err := db.Find(&roles).Error; err != nil {
		log.Printf("Gagal mengambil roles: %v", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("securepassword"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Gagal mengenkripsi password: %v", err)
		return
	}

	user := entity.User{
		Username: "admin_user",
		Password: string(hashedPassword),
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