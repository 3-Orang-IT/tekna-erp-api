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

type ModulManagementHandler struct {
	usecase adminUsecase.ModulManagementUsecase
}

func NewModulManagementHandler(r *gin.Engine, uc adminUsecase.ModulManagementUsecase, db *gorm.DB) {
	h := &ModulManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/moduls", h.CreateModul)
	admin.GET("/moduls", h.GetModuls)
	admin.GET("/moduls/:id", h.GetModulByID)
	admin.PUT("/moduls/:id", h.UpdateModul)
	admin.DELETE("/moduls/:id", h.DeleteModul)
}

func (h *ModulManagementHandler) CreateModul(c *gin.Context) {
	var input dto.CreateModulInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	modul := entity.Modul{
		Name: input.Name,
	}

	if err := h.usecase.CreateModul(&modul); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "modul created successfully", "data": modul})
}

func (h *ModulManagementHandler) GetModuls(c *gin.Context) {
	moduls, err := h.usecase.GetModuls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": moduls})
}

func (h *ModulManagementHandler) GetModulByID(c *gin.Context) {
	id := c.Param("id")
	modul, err := h.usecase.GetModulByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": modul})
}

func (h *ModulManagementHandler) UpdateModul(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid modul ID"})
		return
	}

	var input dto.UpdateModulInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	modul := entity.Modul{
		ID:   uint(idUint),
		Name: input.Name,
	}

	if err := h.usecase.UpdateModul(id, &modul); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "modul updated successfully", "data": modul})
}

func (h *ModulManagementHandler) DeleteModul(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteModul(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "modul deleted successfully", "data": gin.H{"id": id}})
}
