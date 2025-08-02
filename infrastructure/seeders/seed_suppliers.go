package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedSuppliers(db *gorm.DB) error {
	// First get a city ID for reference
	var jakartaCity entity.City
	
	if err := db.Where("name = ?", "Jakarta Pusat").First(&jakartaCity).Error; err != nil {
		log.Printf("Error finding Jakarta Pusat city: %v", err)
		return err
	}

	suppliersList := []entity.Supplier{
		{
			Name:          "PT Supplier Utama",
			InvoiceName:   "PT Supplier Utama",
			NPWP:          "01.234.567.8-123.000",
			Address:       "Jl. Kebon Sirih No. 45",
			CityID:        jakartaCity.ID,
			Phone:         "021-5554321",
			Email:         "contact@supplier-utama.com",
			Greeting:      "Dear Sir/Madam",
			ContactPerson: "Ahmad Supriadi",
			ContactPhone:  "081234567890",
			BankAccount:   "BCA 1234567890",
			Type:          "Material",
			LogoFilename:  "",
		},
		{
			Name:          "CV Mitra Sejahtera",
			InvoiceName:   "CV Mitra Sejahtera",
			NPWP:          "02.345.678.9-123.000",
			Address:       "Jl. Imam Bonjol No. 78",
			CityID:        jakartaCity.ID,
			Phone:         "021-5559876",
			Email:         "info@mitrasejahtera.com",
			Greeting:      "Dear Sir/Madam",
			ContactPerson: "Budi Santoso",
			ContactPhone:  "082345678901",
			BankAccount:   "Mandiri 0987654321",
			Type:          "Equipment",
			LogoFilename:  "",
		},
	}

	for _, supplier := range suppliersList {
		var s entity.Supplier
		if err := db.Where("name = ?", supplier.Name).First(&s).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&supplier).Error; err != nil {
				log.Printf("Failed to create supplier %s: %v", supplier.Name, err)
				return err
			}
			log.Printf("Successfully added supplier %s", supplier.Name)
		}
	}
	return nil
}
