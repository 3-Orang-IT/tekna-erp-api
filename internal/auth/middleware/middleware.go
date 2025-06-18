package middleware

import (
	"net/http"
	"strings"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JWTAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth header format"})
            c.Abort()
            return
        }

        tokenStr := parts[1]
        claims, err := utils.ValidateToken(tokenStr)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        var user entity.User
        if err := db.Preload("Role").First(&user, claims.UserID).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            c.Abort()
            return
        }

        c.Set("userID", claims.UserID)
        c.Next()
    }
}
