package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedMenus(db *gorm.DB) error {
	menusList := []entity.Menu{
		{Name: "Dashboard", ModulID: 1, URL: "/dashboard", Icon: "dashboard", Order: 1},
		{Name: "Users", ModulID: 1, URL: "/users", Icon: "users", Order: 2},
		{Name: "Roles", ModulID: 2, URL: "/roles", Icon: "roles", Order: 1},
	}

	for _, menu := range menusList {
		var m entity.Menu
		if err := db.Where("name = ? AND modul_id = ?", menu.Name, menu.ModulID).First(&m).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&menu).Error; err != nil {
				log.Printf("Failed to create menu %s: %v", menu.Name, err)
				return err
			}
			log.Printf("Successfully added menu %s", menu.Name)
		}
	}
	return nil
}
