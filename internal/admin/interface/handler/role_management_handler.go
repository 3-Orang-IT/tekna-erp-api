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
	admin.GET("/roles/add", h.GetAddRolePage)
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/roles", h.CreateRole)
	admin.GET("/roles", h.GetRoles)
	admin.GET("/roles/:id", h.GetRoleByID)
	admin.GET("/roles/:id/edit", h.GetRoleEditPage)
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

	search := c.DefaultQuery("search", "")

	// Get total count of roles for pagination
	total, err := h.usecase.GetRolesCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	roles, err := h.usecase.GetRoles(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseData []dto.RoleResponse
	for _, role := range roles {
		responseData = append(responseData, dto.RoleResponse{
			ID:   role.ID,
			Name: role.Name,
			Menus: role.Menus,
			CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	response := gin.H{
		"data": responseData,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total_data":  total,
			"total_pages": totalPages,
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

	response := dto.RoleResponse{
		ID:   role.ID,
		Name: role.Name,
		Menus: role.Menus,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
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

func (h *RoleManagementHandler) GetRoleEditPage(c *gin.Context) {
	id := c.Param("id")
	role, err := h.usecase.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch all menus for reference
	var menus []entity.Menu
	if err := h.usecase.GetAllMenus(&menus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data": dto.RoleResponse{
			ID:   role.ID,
			Name: role.Name,
			Menus: role.Menus,
		},
		"reference": gin.H{
			"menus": menus,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *RoleManagementHandler) GetAddRolePage(c *gin.Context) {
	// Fetch all menus for the role creation form
	var menus []entity.Menu
	if err := h.usecase.GetAllMenus(&menus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data": gin.H{
			"menus": menus,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}