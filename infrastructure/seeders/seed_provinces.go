package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedProvinces(db *gorm.DB) error {
	provincesList := []entity.Province{
		{Name: "DKI Jakarta"},
		{Name: "Jawa Barat"},
		{Name: "Jawa Tengah"},
		{Name: "Jawa Timur"},
		{Name: "Bali"},
		{Name: "Sumatera Utara"},
		{Name: "Sumatera Selatan"},
	}

	for _, province := range provincesList {
		var p entity.Province
		if err := db.Where("name = ?", province.Name).First(&p).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&province).Error; err != nil {
				log.Printf("Failed to create province %s: %v", province.Name, err)
				return err
			}
			log.Printf("Successfully added province %s", province.Name)
		}
	}
	return nil
}
