package adminHandler

import (
	"net/http"
	"strconv"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/middleware"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleManagementHandler struct {
	usecase adminUsecase.RoleManagementUsecase
}

func NewRoleManagementHandler(r *gin.Engine, uc adminUsecase.RoleManagementUsecase, db *gorm.DB) {
	h := &RoleManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/roles", h.CreateRole)
	admin.GET("/roles", h.GetRoles)
	admin.GET("/roles/:id", h.GetRoleByID)
	admin.PUT("/roles/:id", h.UpdateRole)
	admin.DELETE("/roles/:id", h.DeleteRole)
}

func (h *RoleManagementHandler) CreateRole(c *gin.Context) {
	var input dto.CreateRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var menus []entity.Menu
	for _, menuID := range input.MenuIDs {
		menus = append(menus, entity.Menu{ID: menuID})
	}

	role := entity.Role{
		Name:  input.Name,
		Menus: menus,
	}

	if err := h.usecase.CreateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "role created successfully", "data": role})
}

func (h *RoleManagementHandler) GetRoles(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	roles, err := h.usecase.GetRoles(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data": roles,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *RoleManagementHandler) GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	role, err := h.usecase.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

func (h *RoleManagementHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role ID"})
		return
	}

	var input dto.UpdateRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var menus []entity.Menu
	for _, menuID := range input.MenuIDs {
		menus = append(menus, entity.Menu{ID: menuID})
	}

	role := entity.Role{
		ID:    uint(idUint),
		Name:  input.Name,
		Menus: menus,
	}

	if err := h.usecase.UpdateRole(id, &role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role updated successfully", "data": role})
}

func (h *RoleManagementHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteRole(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role deleted successfully", "data": gin.H{"id": id}})
}
