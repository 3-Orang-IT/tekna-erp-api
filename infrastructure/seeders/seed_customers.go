package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedCustomers(db *gorm.DB) error {
	customers := []entity.Customer{
		{
			UserID:            1,
			AreaID:            1,
			CityID:            1,
			Name:              "PT Sukses Selalu",
			Code:              "CUST-1",
			InvoiceName:       "PT Sukses Selalu",
			Address:           "Jl. Merdeka No. 1",
			Phone:             "081234567890",
			Email:             "info@suksesselalu.com",
			Tax:               "123456789",
			Greeting:          "Dear Customer",
			ContactPersonName: "John Doe",
			ContactPhone:      "081234567891",
			Segment:           "Retail",
			Type:              "Distributor",
			NPWP:              "987654321",
			Status:            "Active",
			BEName:            "Jane Smith",
			ProcurementType:   "Direct",
			MarketingName:     "Mark Johnson",
			Note:              "Top customer",
			PaymentTerm:       "30 days",
			Level:             "Gold",
		},
		{
			UserID:            1,
			AreaID:            2,
			CityID:            2,
			Name:              "PT Makmur Jaya",
			Code:              "CUST-002",
			InvoiceName:       "PT Makmur Jaya",
			Address:           "Jl. Kebahagiaan No. 2",
			Phone:             "081234567892",
			Email:             "info@makmurjaya.com",
			Tax:               "2233445566",
			Greeting:          "Dear Partner",
			ContactPersonName: "Alice Brown",
			ContactPhone:      "081234567893",
			Segment:           "Wholesale",
			Type:              "Retailer",
			NPWP:              "1122334455",
			Status:            "Active",
			BEName:            "Bob White",
			ProcurementType:   "Indirect",
			MarketingName:     "Charlie Green",
			Note:              "Preferred partner",
			PaymentTerm:       "45 days",
			Level:             "Silver",
		},
	}

	for _, customer := range customers {
		if err := db.FirstOrCreate(&customer, entity.Customer{ID: customer.ID}).Error; err != nil {
			log.Printf("Error seeding customer with ID %d: %v", customer.ID, err)
			return err
		}
	}

	return nil
}
