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

type BusinessUnitManagementHandler struct {
	usecase adminUsecase.BusinessUnitManagementUsecase
}

func NewBusinessUnitManagementHandler(r *gin.Engine, uc adminUsecase.BusinessUnitManagementUsecase, db *gorm.DB) {
	h := &BusinessUnitManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/business-units", h.CreateBusinessUnit)
	admin.GET("/business-units", h.GetBusinessUnits)
	admin.GET("/business-units/:id", h.GetBusinessUnitByID)
	admin.GET("/business-units/:id/edit", h.GetEditBusinessUnitPage)
	admin.PUT("/business-units/:id", h.UpdateBusinessUnit)
	admin.DELETE("/business-units/:id", h.DeleteBusinessUnit)
}

func (h *BusinessUnitManagementHandler) CreateBusinessUnit(c *gin.Context) {
	var input dto.CreateBusinessUnitInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	businessUnit := entity.BusinessUnit{
		Name: input.Name,
	}

	if err := h.usecase.CreateBusinessUnit(&businessUnit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "business unit created successfully", "data": businessUnit})
}

func (h *BusinessUnitManagementHandler) GetBusinessUnits(c *gin.Context) {
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

	businessUnits, err := h.usecase.GetBusinessUnits(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data": businessUnits,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *BusinessUnitManagementHandler) GetBusinessUnitByID(c *gin.Context) {
	id := c.Param("id")
	businessUnit, err := h.usecase.GetBusinessUnitByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": businessUnit})
}

func (h *BusinessUnitManagementHandler) UpdateBusinessUnit(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateBusinessUnitInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	businessUnit := entity.BusinessUnit{
		Name: input.Name,
	}

	if err := h.usecase.UpdateBusinessUnit(id, &businessUnit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "business unit updated successfully"})
}

func (h *BusinessUnitManagementHandler) DeleteBusinessUnit(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteBusinessUnit(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "business unit deleted successfully"})
}

func (h *BusinessUnitManagementHandler) GetEditBusinessUnitPage(c *gin.Context) {
	id := c.Param("id")
	businessUnit, err := h.usecase.GetBusinessUnitByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": businessUnit})
}
