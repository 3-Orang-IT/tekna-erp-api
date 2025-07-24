package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedCompanies(db *gorm.DB) error {
	// First get a city ID for reference
	var jakartaCity entity.City
	
	if err := db.Where("name = ?", "Jakarta Pusat").First(&jakartaCity).Error; err != nil {
		log.Printf("Error finding Jakarta Pusat city: %v", err)
		return err
	}

	companiesList := []entity.Company{
		{
			Name:             "PT Tekna Indonesia",
			Address:          "Jl. Sudirman No. 123",
			CityID:           jakartaCity.ID,
			Phone:            "021-5551234",
			Fax:              "021-5551235",
			Email:            "contact@tekna.id",
			StartHour:        "08:00",
			EndHour:          "17:00",
			Latitude:         -6.2088,
			Longitude:        106.8456,
			TotalShares:      10000,
			AnnualLeaveQuota: 12,
		},
		{
			Name:             "PT Tekna Solutions",
			Address:          "Jl. Gatot Subroto No. 456",
			CityID:           jakartaCity.ID,
			Phone:            "021-5557890",
			Fax:              "021-5557891",
			Email:            "info@tekna-solutions.id",
			StartHour:        "09:00",
			EndHour:          "18:00",
			Latitude:         -6.2256,
			Longitude:        106.8025,
			TotalShares:      5000,
			AnnualLeaveQuota: 14,
		},
	}

	for _, company := range companiesList {
		var c entity.Company
		if err := db.Where("name = ?", company.Name).First(&c).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&company).Error; err != nil {
				log.Printf("Failed to create company %s: %v", company.Name, err)
				return err
			}
			log.Printf("Successfully added company %s", company.Name)
		}
	}
	return nil
}
