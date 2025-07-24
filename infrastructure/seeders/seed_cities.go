package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedCities(db *gorm.DB) error {
	// First get the province IDs for reference
	var jakartaProvince, westJavaProvince, centralJavaProvince entity.Province
	
	// Get Jakarta province
	if err := db.Where("name = ?", "DKI Jakarta").First(&jakartaProvince).Error; err != nil {
		log.Printf("Error finding DKI Jakarta province: %v", err)
		return err
	}
	
	// Get Jawa Barat province
	if err := db.Where("name = ?", "Jawa Barat").First(&westJavaProvince).Error; err != nil {
		log.Printf("Error finding Jawa Barat province: %v", err)
		return err
	}
	
	// Get Jawa Tengah province
	if err := db.Where("name = ?", "Jawa Tengah").First(&centralJavaProvince).Error; err != nil {
		log.Printf("Error finding Jawa Tengah province: %v", err)
		return err
	}

	// Define city list with their respective province IDs
	citiesList := []entity.City{
		{Name: "Jakarta Pusat", ProvinceID: &jakartaProvince.ID},
		{Name: "Jakarta Selatan", ProvinceID: &jakartaProvince.ID},
		{Name: "Jakarta Timur", ProvinceID: &jakartaProvince.ID},
		{Name: "Jakarta Barat", ProvinceID: &jakartaProvince.ID},
		{Name: "Jakarta Utara", ProvinceID: &jakartaProvince.ID},
		{Name: "Bandung", ProvinceID: &westJavaProvince.ID},
		{Name: "Bogor", ProvinceID: &westJavaProvince.ID},
		{Name: "Bekasi", ProvinceID: &westJavaProvince.ID},
		{Name: "Depok", ProvinceID: &westJavaProvince.ID},
		{Name: "Semarang", ProvinceID: &centralJavaProvince.ID},
		{Name: "Solo", ProvinceID: &centralJavaProvince.ID},
		{Name: "Yogyakarta", ProvinceID: &centralJavaProvince.ID},
	}

	for _, city := range citiesList {
		var c entity.City
		if err := db.Where("name = ?", city.Name).First(&c).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&city).Error; err != nil {
				log.Printf("Failed to create city %s: %v", city.Name, err)
				return err
			}
			log.Printf("Successfully added city %s", city.Name)
		}
	}
	return nil
}
