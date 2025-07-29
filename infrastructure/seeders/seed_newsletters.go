package seeders

import (
	"time"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedNewsletters(db *gorm.DB) error {
	var count int64
	db.Model(&entity.Newsletter{}).Count(&count)
	if count > 0 {
		return nil // Newsletters already seeded
	}

	newsletters := []entity.Newsletter{
		{
			Type:        "Announcement",
			Title:       "Welcome to Company Portal",
			Description: "Welcome to our new company portal. Please explore the new features!",
			File:        "welcome_announcement.pdf",
			ValidFrom:   "2023-07-01",
			Status:      "Active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Type:        "News",
			Title:       "New Product Launch",
			Description: "We are excited to announce our new product launch next month!",
			File:        "product_launch.pdf",
			ValidFrom:   "2023-07-15",
			Status:      "Active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Type:        "Promo",
			Title:       "Summer Sale Promotion",
			Description: "Enjoy our summer promotion with up to 50% discount!",
			File:        "summer_promo.pdf",
			ValidFrom:   "2023-08-01",
			Status:      "Draft",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	return db.Create(&newsletters).Error
}
