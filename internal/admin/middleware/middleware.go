package middleware

import (
	"net/http"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoleMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID missing in context"})
			c.Abort()
			return
		}

		var user entity.User
		if err := db.Preload("Role").First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		isAdmin := false
		for _, role := range user.Role {
			if role.ID == 1 {
				isAdmin = true
				break
			}
		}

		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: admin role required"})
			c.Abort()
			return
		}

		c.Next()
	}
}
