package adminHandler

import (
	"net/http"
	"strconv"
	"strings"

	"errors"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/middleware"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AreaManagementHandler struct {
	usecase adminUsecase.AreaManagementUsecase
}

func NewAreaManagementHandler(r *gin.Engine, uc adminUsecase.AreaManagementUsecase, db *gorm.DB) {
	h := &AreaManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/areas", h.CreateArea)
	admin.GET("/areas", h.GetAreas)
	admin.GET("/areas/:id", h.GetAreaByID)
	admin.GET("/areas/:id/edit", h.GetEditAreaPage)
	admin.PUT("/areas/:id", h.UpdateArea)
	admin.DELETE("/areas/:id", h.DeleteArea)
}

func (h *AreaManagementHandler) CreateArea(c *gin.Context) {
	var input dto.CreateAreaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Make sure ID is not set
	area := entity.Area{
		Name: input.Name,
	}

	if err := h.usecase.CreateArea(&area); err != nil {
		// Check if the error is related to duplicate entry
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			c.JSON(http.StatusConflict, gin.H{"error": "An area with this name already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create area"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "area created successfully", "data": area})
}

func (h *AreaManagementHandler) GetAreas(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.DefaultQuery("search", "")

	// Get total count of areas for pagination
	total, err := h.usecase.GetAreasCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get areas count"})
		return
	}

	// Calculate total pages
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	areas, err := h.usecase.GetAreas(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch areas"})
		return
	}

	// Map areas to the desired response format using AreaResponse DTO
	var responseData []dto.AreaResponse
	for _, area := range areas {
		// Extract employee names or IDs to match the DTO structure
		var employeeNames []string
		if area.Employees != nil {
			for _, employee := range area.Employees {
				if employee.User.Name != "" {
					employeeNames = append(employeeNames, employee.User.Name)
				}
			}
		}
		
		responseData = append(responseData, dto.AreaResponse{
			ID:        area.ID,
			Name:      area.Name,
			Employees: employeeNames,
			CreatedAt: area.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt: area.UpdatedAt.Format("02-01-2006 15:04:05"),
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

func (h *AreaManagementHandler) GetAreaByID(c *gin.Context) {
	id := c.Param("id")
	area, err := h.usecase.GetAreaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Area not found"})
		return
	}

	// Extract employee names to match the DTO structure
	var employeeNames []string
	if area.Employees != nil {
		for _, employee := range area.Employees {
			if employee.User.Name != "" {
				employeeNames = append(employeeNames, employee.User.Name)
			}
		}
	}

	// Map to AreaResponse DTO for clean response
	response := dto.AreaResponse{
		ID:        area.ID,
		Name:      area.Name,
		Employees: employeeNames,
		CreatedAt: area.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: area.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *AreaManagementHandler) UpdateArea(c *gin.Context) {
	id := c.Param("id")
	
	// Validate if area exists
	_, err := h.usecase.GetAreaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Area not found"})
		return
	}

	var input dto.UpdateAreaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	area := entity.Area{
		Name: input.Name,
	}

	if err := h.usecase.UpdateArea(id, &area); err != nil {
		// Check if the error is related to duplicate entry
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "duplicate") {
			c.JSON(http.StatusConflict, gin.H{"error": "An area with this name already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update area"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "area updated successfully"})
}

func (h *AreaManagementHandler) DeleteArea(c *gin.Context) {
	id := c.Param("id")
	
	// Validate if area exists
	_, err := h.usecase.GetAreaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Area not found"})
		return
	}

	if err := h.usecase.DeleteArea(id); err != nil {
		// Special handling for the case when it's the last area
		if strings.Contains(err.Error(), "last remaining area") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		// Handle any other errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete area: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Area deleted successfully. Any customers using this area have been reassigned to another area, and employee associations have been removed."})
}

// GetEditAreaPage returns the data needed for the edit area page
func (h *AreaManagementHandler) GetEditAreaPage(c *gin.Context) {
	id := c.Param("id")
	area, err := h.usecase.GetAreaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Area not found"})
		return
	}

	// Extract employee names to match the DTO structure
	var employeeNames []string
	if area.Employees != nil {
		for _, employee := range area.Employees {
			if employee.User.Name != "" {
				employeeNames = append(employeeNames, employee.User.Name)
			}
		}
	}

	response := dto.AreaResponse{
		ID:        area.ID,
		Name:      area.Name,
		Employees: employeeNames,
		CreatedAt: area.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: area.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
