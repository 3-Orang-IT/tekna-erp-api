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

type UnitOfMeasureManagementHandler struct {
	usecase adminUsecase.UnitOfMeasureManagementUsecase
}

func NewUnitOfMeasureManagementHandler(r *gin.Engine, uc adminUsecase.UnitOfMeasureManagementUsecase, db *gorm.DB) {
	h := &UnitOfMeasureManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/unit-of-measures", h.CreateUnitOfMeasure)
	admin.GET("/unit-of-measures", h.GetUnitOfMeasures)
	admin.GET("/unit-of-measures/:id", h.GetUnitOfMeasureByID)
	admin.GET("/unit-of-measures/:id/edit", h.GetEditUnitOfMeasurePage)
	admin.PUT("/unit-of-measures/:id", h.UpdateUnitOfMeasure)
	admin.DELETE("/unit-of-measures/:id", h.DeleteUnitOfMeasure)
}

func (h *UnitOfMeasureManagementHandler) CreateUnitOfMeasure(c *gin.Context) {
	var input dto.CreateUnitOfMeasureInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unitOfMeasure := entity.UnitOfMeasure{
		Name:         input.Name,
		Abbreviation: input.Abbreviation,
	}

	if err := h.usecase.CreateUnitOfMeasure(&unitOfMeasure); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "unit of measure created successfully", "data": unitOfMeasure})
}

func (h *UnitOfMeasureManagementHandler) GetUnitOfMeasures(c *gin.Context) {
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

	// Get total count of units of measure for pagination
	total, err := h.usecase.GetUnitOfMeasuresCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	unitOfMeasures, err := h.usecase.GetUnitOfMeasures(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data": unitOfMeasures,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total_data":  total,
			"total_pages": totalPages,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *UnitOfMeasureManagementHandler) GetUnitOfMeasureByID(c *gin.Context) {
	id := c.Param("id")
	unitOfMeasure, err := h.usecase.GetUnitOfMeasureByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": unitOfMeasure})
}

func (h *UnitOfMeasureManagementHandler) UpdateUnitOfMeasure(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateUnitOfMeasureInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	unitOfMeasure := entity.UnitOfMeasure{
		Name:         input.Name,
		Abbreviation: input.Abbreviation,
	}

	if err := h.usecase.UpdateUnitOfMeasure(id, &unitOfMeasure); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "unit of measure updated successfully"})
}

func (h *UnitOfMeasureManagementHandler) DeleteUnitOfMeasure(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteUnitOfMeasure(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "unit of measure deleted successfully"})
}

func (h *UnitOfMeasureManagementHandler) GetEditUnitOfMeasurePage(c *gin.Context) {
	id := c.Param("id")
	unitOfMeasure, err := h.usecase.GetUnitOfMeasureByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": unitOfMeasure})
}
