package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedModuls(db *gorm.DB) error {
	modulsList := []entity.Modul{
		{Name: "User Management"},
		{Name: "Role Management"},
		{Name: "Menu Management"},
	}

	for _, modul := range modulsList {
		var m entity.Modul
		if err := db.Where("name = ?", modul.Name).First(&m).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&modul).Error; err != nil {
				log.Printf("Failed to create modul %s: %v", modul.Name, err)
				return err
			}
			log.Printf("Successfully added modul %s", modul.Name)
		}
	}
	return nil
}
