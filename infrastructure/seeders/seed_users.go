package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	var roles []entity.Role
	if err := db.Find(&roles).Error; err != nil {
		log.Printf("Failed to fetch roles: %v", err)
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("securepassword"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
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
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == gorm.ErrRecordNotFound {
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Failed to create user %s: %v", user.Username, err)
			return err
		}
		log.Printf("Successfully added user %s with all roles", user.Username)
	}
	return nil
}
