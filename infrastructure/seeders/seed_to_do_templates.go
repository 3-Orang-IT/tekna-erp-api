package seeders

import (
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedToDoTemplates(tx *gorm.DB) error {
	var count int64
	tx.Model(&entity.ToDoTemplate{}).Count(&count)
	if count > 0 {
		return nil // To Do Templates already seeded
	}

	// Get job positions to associate with to-do templates
	var jobPositions []entity.JobPosition
	if err := tx.Limit(3).Find(&jobPositions).Error; err != nil {
		return err
	}

	// If we have job positions, create some to-do templates
	if len(jobPositions) > 0 {
		toDoTemplates := []entity.ToDoTemplate{
			{
				JobPositionID: jobPositions[0].ID,
				Activity:      "Daily standup meeting",
				Priority:      1,
				OrderNumber:   1,
			},
			{
				JobPositionID: jobPositions[0].ID,
				Activity:      "Code review",
				Priority:      2,
				OrderNumber:   2,
			},
			{
				JobPositionID: jobPositions[1 % len(jobPositions)].ID,
				Activity:      "Weekly report submission",
				Priority:      1,
				OrderNumber:   1,
			},
			{
				JobPositionID: jobPositions[1 % len(jobPositions)].ID,
				Activity:      "Client meeting preparation",
				Priority:      2,
				OrderNumber:   2,
			},
			{
				JobPositionID: jobPositions[2 % len(jobPositions)].ID,
				Activity:      "Monthly performance review",
				Priority:      1,
				OrderNumber:   1,
			},
		}

		for _, toDoTemplate := range toDoTemplates {
			if err := tx.Create(&toDoTemplate).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
