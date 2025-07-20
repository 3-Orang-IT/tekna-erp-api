package adminHandler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/middleware"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserManagementHandler struct {
	usecase adminUsecase.UserManagementUsecase
}

func NewUserManagementHandler(r *gin.Engine, uc adminUsecase.UserManagementUsecase,  db *gorm.DB) {
	h := &UserManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/users", h.CreateUser)
	admin.GET("/users", h.GetUsers)
	admin.GET("/users/:id", h.GetUserByID)
	admin.PUT("/users/:id", h.UpdateUser)
	admin.DELETE("/users/:id", h.DeleteUser)
}

func validateUser(user *entity.User) error {
	if strings.TrimSpace(user.Username) == "" {
		return fmt.Errorf("username is required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return fmt.Errorf("email is required")
	}
	if strings.TrimSpace(user.Password) == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

func (h *UserManagementHandler) CreateUser(c *gin.Context) {
	var input dto.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{
		Username:        input.Username,
		Password:        input.Password,
		Name:            input.Name,
		Email:           input.Email,
		Telp:            input.Telp,
		PhotoProfileURL: input.PhotoProfileURL,
		Status:          input.Status,
	}

	for _, roleID := range input.RoleIDs {
		user.Role = append(user.Role, entity.Role{ID: roleID})
	}

	if err := validateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "data": user})
}

func (h *UserManagementHandler) GetUsers(c *gin.Context) {
	users, err := h.usecase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *UserManagementHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *UserManagementHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.UpdateUser(id, &user); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "user ID not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully", "data": user})
}

func (h *UserManagementHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteUser(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "user ID not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
		"data": gin.H{"id": id},
	})
}