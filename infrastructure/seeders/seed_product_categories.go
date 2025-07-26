package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedProductCategories(db *gorm.DB) error {
	categoriesList := []entity.ProductCategory{
		{Name: "Electronics"},
		{Name: "Furniture"},
		{Name: "Clothing"},
		{Name: "Food & Beverage"},
		{Name: "Stationery"},
	}

	for _, category := range categoriesList {
		var c entity.ProductCategory
		if err := db.Where("name = ?", category.Name).First(&c).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Failed to create product category %s: %v", category.Name, err)
				return err
			}
			log.Printf("Successfully added product category %s", category.Name)
		}
	}
	return nil
}
