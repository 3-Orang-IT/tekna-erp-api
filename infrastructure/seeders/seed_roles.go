package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	rolesList := []entity.Role{
		{Name: "Admin"},
		{Name: "Team Support"},
		{Name: "HRD"},
		{Name: "Supplier"},
	}

	for _, role := range rolesList {
		var r entity.Role
		if err := db.Where("name = ?", role.Name).First(&r).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&role).Error; err != nil {
				log.Printf("Failed to create role %s: %v", role.Name, err)
				return err
			}
			log.Printf("Successfully added role %s", role.Name)
		}
	}
	return nil
}
