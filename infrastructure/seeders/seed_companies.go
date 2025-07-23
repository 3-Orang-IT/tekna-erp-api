package seeders

import (
    "log"
    "time"

    "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
    "gorm.io/gorm"
)

func SeedCompanies(db *gorm.DB) error {
    company := entity.Company{
        ID:               "9bd50663-34f1-412d-9e28-a6ae23be5576",
        Name:             "PT. Mitra Karya Analitika",
        Address:          "Jl. Klipang Raya, Ruko Amsterdam No.90, Semarang",
        CityID:           "3374",
        ProvinceID:       "33",
        Phone:            "024-76412142",
        Fax:              "",
        Email:            "support@mikacares.com",
        StartHour:        "08:00:00",
        EndHour:          "17:00:00",
        Latitude:         -7.049414065937655,
        Longitude:        110.48074422452338,
        TotalShares:      500,
        AnnualLeaveQuota: 12,
        CreatedAt:        time.Date(2020, 9, 25, 9, 59, 52, 0, time.UTC),
        UpdatedAt:        time.Date(2021, 5, 23, 13, 16, 24, 0, time.UTC),
    }

    var existingCompany entity.Company
    if err := db.Where("id = ?", company.ID).First(&existingCompany).Error; err == gorm.ErrRecordNotFound {
        if err := db.Create(&company).Error; err != nil {
            log.Printf("Failed to create company %s: %v", company.Name, err)
            return err
        }
        log.Printf("Successfully added company %s", company.Name)
    } else {
        log.Printf("Company %s already exists", company.Name)
    }

    return nil
}