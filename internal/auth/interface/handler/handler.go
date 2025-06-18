package handler

import (
	"fmt"
	"net/http"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/middleware"
	usecase "github.com/3-Orang-IT/tekna-erp-api/internal/auth/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
    usecase usecase.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, uc usecase.AuthUsecase, db *gorm.DB) {
    h := &AuthHandler{uc}
    api := r.Group("/api/v1")
    api.POST("/auth/register", h.Register)
    api.POST("/auth/login", h.Login)

    protected := api.Group("/")
    protected.Use(middleware.JWTAuthMiddleware(db))
    protected.GET("/menus", h.GetMenus)
}

func (h *AuthHandler) Register(c *gin.Context) {
    var user entity.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.usecase.Register(&user); err != nil {
        if err.Error() == "email already registered" {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var input struct {
        Username    string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.usecase.Login(input.Username, input.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
        return
    }

    roles := []gin.H{}
    for _, role := range user.Role {
        roles = append(roles, gin.H{
            "id":   role.ID,
            "name": role.Name,
        })
    }

    c.JSON(http.StatusOK, gin.H{
        "data" : gin.H{
            "token": token,
            "user":  gin.H{
                "id":    user.ID,
                "name":  user.Name,
                "email": user.Email,
                "role": roles,
            },
        },
    })
}

func (h *AuthHandler) GetMenus(c *gin.Context) {
    userIdInterface, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
        return
    }

    userID, ok := userIdInterface.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
        return
    }

    menus, err := h.usecase.GetMenus(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, menus)
}



func parseUint(str string) uint {
    var i uint
    fmt.Sscanf(str, "%d", &i)
    return i
}
