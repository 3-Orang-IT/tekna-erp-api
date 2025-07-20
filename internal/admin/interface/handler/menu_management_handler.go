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

type MenuManagementHandler struct {
	usecase adminUsecase.MenuManagementUsecase
}

func NewMenuManagementHandler(r *gin.Engine, uc adminUsecase.MenuManagementUsecase, db *gorm.DB) {
	h := &MenuManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/menus", h.CreateMenu)
	admin.GET("/menus", h.GetMenus)
	admin.GET("/menus/:id", h.GetMenuByID)
	admin.PUT("/menus/:id", h.UpdateMenu)
	admin.DELETE("/menus/:id", h.DeleteMenu)
}

func (h *MenuManagementHandler) CreateMenu(c *gin.Context) {
	var input dto.CreateMenuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menu := entity.Menu{
		Name:     input.Name,
		URL:      input.URL,
		Icon:     input.Icon,
		Order:    input.Order,
		ParentID: input.ParentID,
		ModulID:  input.ModulID,
	}

	if err := h.usecase.CreateMenu(&menu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "menu created successfully", "data": menu})
}

func (h *MenuManagementHandler) GetMenus(c *gin.Context) {
	menus, err := h.usecase.GetMenus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": menus})
}

func (h *MenuManagementHandler) GetMenuByID(c *gin.Context) {
	id := c.Param("id")
	menu, err := h.usecase.GetMenuByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": menu})
}

func (h *MenuManagementHandler) UpdateMenu(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid menu ID"})
		return
	}

	var input dto.UpdateMenuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menu := entity.Menu{
		ID:       uint(idUint),
		Name:     input.Name,
		URL:      input.URL,
		Icon:     input.Icon,
		Order:    input.Order,
		ParentID: input.ParentID,
		ModulID:  input.ModulID,
	}

	if err := h.usecase.UpdateMenu(id, &menu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "menu updated successfully", "data": menu})
}

func (h *MenuManagementHandler) DeleteMenu(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteMenu(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "menu deleted successfully", "data": gin.H{"id": id}})
}
