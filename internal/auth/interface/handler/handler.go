package handler

import (
	"fmt"
	"net/http"

	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/domain/entity"
	usecase "github.com/3-Orang-IT/tekna-erp-api/internal/auth/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    usecase usecase.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, uc usecase.AuthUsecase) {
    h := &AuthHandler{uc}
    api := r.Group("/api/v1/auth")
    api.POST("/register", h.Register)
    api.POST("/login", h.Login)
    api.GET("/menus/:role_id", h.GetMenus)
}

func (h *AuthHandler) Register(c *gin.Context) {
    var user entity.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.usecase.Register(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.usecase.Login(input.Email, input.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *AuthHandler) GetMenus(c *gin.Context) {
    // ambil role id dari path
    roleID := c.Param("role_id")
    menus, err := h.usecase.GetMenus(parseUint(roleID))
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
