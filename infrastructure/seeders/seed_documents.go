package seeders

import (
	"time"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedDocuments(db *gorm.DB) error {
	var count int64
	db.Model(&entity.Document{}).Count(&count)
	if count > 0 {
		return nil // Documents already seeded
	}

	// Check if we have categories and users
	var categoryCount int64
	db.Model(&entity.DocumentCategory{}).Count(&categoryCount)
	if categoryCount == 0 {
		return nil // No categories yet, skip seeding documents
	}

	var userCount int64
	db.Model(&entity.User{}).Count(&userCount)
	if userCount == 0 {
		return nil // No users yet, skip seeding documents
	}

	// Get admin user ID
	var adminUser entity.User
	if err := db.Where("username = ?", "admin_user").First(&adminUser).Error; err != nil {
		return nil // Admin user not found, skip seeding
	}

	// Get first document category
	var firstCategory entity.DocumentCategory
	if err := db.First(&firstCategory).Error; err != nil {
		return nil // Category not found, skip seeding
	}

	documents := []entity.Document{
		{
			DocumentCategoryID: firstCategory.ID,
			Name:               "Company Policy Document",
			UserID:             adminUser.ID,
			FilePath:           "uploads/documents/sample_policy.pdf",
			Description:        "This document outlines the company policies and procedures.",
			IsPublished:        true,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		},
		{
			DocumentCategoryID: firstCategory.ID,
			Name:               "Employee Handbook",
			UserID:             adminUser.ID,
			FilePath:           "uploads/documents/employee_handbook.pdf",
			Description:        "A comprehensive guide for all employees.",
			IsPublished:        true,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		},
	}

	return db.Create(&documents).Error
}
