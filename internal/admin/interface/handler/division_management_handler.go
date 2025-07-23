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

type DivisionManagementHandler struct {
	usecase adminUsecase.DivisionManagementUsecase
}

func NewDivisionManagementHandler(r *gin.Engine, uc adminUsecase.DivisionManagementUsecase, db *gorm.DB) {
	h := &DivisionManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/divisions", h.CreateDivision)
	admin.GET("/divisions", h.GetDivisions)
	admin.GET("/divisions/:id", h.GetDivisionByID)
	admin.PUT("/divisions/:id", h.UpdateDivision)
	admin.DELETE("/divisions/:id", h.DeleteDivision)
}

func (h *DivisionManagementHandler) CreateDivision(c *gin.Context) {
	var input dto.CreateDivisionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	division := entity.Division{
		Name: input.Name,
	}

	if err := h.usecase.CreateDivision(&division); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "division created successfully", "data": division})
}

func (h *DivisionManagementHandler) GetDivisions(c *gin.Context) {
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

	divisions, err := h.usecase.GetDivisions(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": divisions, "page": page, "limit": limit})
}

func (h *DivisionManagementHandler) GetDivisionByID(c *gin.Context) {
	id := c.Param("id")
	division, err := h.usecase.GetDivisionByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "division not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": division})
}

func (h *DivisionManagementHandler) UpdateDivision(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateDivisionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	division := entity.Division{
		Name: input.Name,
	}

	if err := h.usecase.UpdateDivision(id, &division); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "division not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "division updated successfully", "data": division})
}

func (h *DivisionManagementHandler) DeleteDivision(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteDivision(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "division not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "division deleted successfully", "data": gin.H{"id": id}})
}
