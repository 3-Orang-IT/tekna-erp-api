package seeders

import (
	"log"
	"time"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	log.Println("Memulai seeding database...")
	SeedRoles(db)
	SeedCompany(db)
	SeedDivisions(db)
	SeedDefaultUser(db)

	log.Println("Seeding database selesai")
}

func SeedDivisions(db *gorm.DB) {
	log.Println("Seeding divisions...")
	divisions := []entity.Division{
		{Name: "Team Support"},
		{Name: "Management"},
		{Name: "Marketing"},
		{Name: "Magang"},
	}
	for _, division := range divisions {
		var d entity.Division
		result := db.Where("name = ?", division.Name).First(&d)
		if result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&division).Error; err != nil {
				log.Printf("Gagal membuat division %s: %v", division.Name, err)
			} else {
				log.Printf("✓ Berhasil menambahkan division: %s", division.Name)
			}
		} else {
			log.Printf("✓ Division %s sudah ada", division.Name)
		}
	}
}

func SeedRoles(db *gorm.DB) {
	log.Println("Seeding roles...")

	rolesList := []entity.Role{
		{Name: "Admin"},
		{Name: "Team Support"},
		{Name: "HRD"},
		{Name: "Supplier"},
	}

	for _, role := range rolesList {
		var r entity.Role
		result := db.Where("name = ?", role.Name).First(&r)
		if result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&role).Error; err != nil {
				log.Printf("Gagal membuat role %s: %v", role.Name, err)
			} else {
				log.Printf("✓ Berhasil menambahkan role: %s", role.Name)
			}
		} else {
			log.Printf("✓ Role %s sudah ada", role.Name)
		}
	}
}

func SeedCompany(db *gorm.DB) {
	log.Println("Seeding company...")

	company := entity.Company{
		ID:               "9bd50663-34f1-412d-9e28-a6ae23be5576",
		Name:             "PT. Mitra Karya Analitika",
		Address:          "Jl. Klipang Raya, Ruko Amsterdam No.90, Semarang",
		CityID:           "Kota Semarang",
		ProvinceID:       "Jawa Tengah",
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
	result := db.Where("id = ?", company.ID).First(&existingCompany)
	if result.Error == gorm.ErrRecordNotFound {
		if err := db.Create(&company).Error; err != nil {
			log.Printf("Gagal membuat company %s: %v", company.Name, err)
		} else {
			log.Printf("✓ Berhasil menambahkan company: %s", company.Name)
		}
	} else {
		log.Printf("✓ Company %s sudah ada", company.Name)
	}
}

func SeedDefaultUser(db *gorm.DB) {
	log.Println("Seeding default user...")

	// Ambil semua roles
	var roles []entity.Role
	if err := db.Find(&roles).Error; err != nil {
		log.Printf("Gagal mengambil roles: %v", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("securepassword"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Gagal mengenkripsi password: %v", err)
		return
	}

	user := entity.User{
		Username: "admin_user",
		Password: string(hashedPassword),
		Name:     "Admin User",
		Email:    "admin@example.com",
		Telp:     "123456789",
		Status:   "active",
		Role:     roles,
	}

	var existingUser entity.User
	result := db.Where("username = ?", user.Username).First(&existingUser)
	if result.Error == gorm.ErrRecordNotFound {
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Gagal membuat user %s: %v", user.Username, err)
		} else {
			log.Printf("✓ Berhasil menambahkan user: %s dengan semua role", user.Username)
		}
	} else {
		log.Printf("✓ User %s sudah ada", user.Username)
	}
}