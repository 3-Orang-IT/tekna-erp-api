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

type TravelCostManagementHandler struct {
	usecase adminUsecase.TravelCostManagementUsecase
}

func NewTravelCostManagementHandler(r *gin.Engine, uc adminUsecase.TravelCostManagementUsecase, db *gorm.DB) {
	h := &TravelCostManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/travel-costs", h.CreateTravelCost)
	admin.GET("/travel-costs", h.GetTravelCosts)
	admin.GET("/travel-costs/add", h.GetAddTravelCostPage)
	admin.GET("/travel-costs/:id", h.GetTravelCostByID)
	admin.GET("/travel-costs/:id/edit", h.GetEditTravelCostPage)
	admin.PUT("/travel-costs/:id", h.UpdateTravelCost)
	admin.DELETE("/travel-costs/:id", h.DeleteTravelCost)
}

func (h *TravelCostManagementHandler) CreateTravelCost(c *gin.Context) {
	var input dto.CreateTravelCostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travelCost := entity.TravelCost{
		Name:  input.Name,
		Unit:  input.Unit,
		Price: input.Price,
	}
	
	// Generate an auto-incrementing travel cost code
	lastTravelCost, err := h.usecase.GetLastTravelCost()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate travel cost code"})
		return
	}
	travelCost.Code = "TRC-" + strconv.Itoa(int(lastTravelCost.ID)+1)

	if err := h.usecase.CreateTravelCost(&travelCost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "travel cost created successfully", "data": travelCost})
}

func (h *TravelCostManagementHandler) GetTravelCosts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.DefaultQuery("search", "")

	// Get total count of travel costs for pagination
	total, err := h.usecase.GetTravelCostsCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	travelCosts, err := h.usecase.GetTravelCosts(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map travel costs to the desired response format using TravelCostResponse DTO
	var responseData []dto.TravelCostResponse
	for _, travelCost := range travelCosts {
		responseData = append(responseData, dto.TravelCostResponse{
			ID:        travelCost.ID,
			Name:      travelCost.Name,
			Code:      travelCost.Code,
			Unit:      travelCost.Unit,
			Price:     travelCost.Price,
			CreatedAt: travelCost.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt: travelCost.UpdatedAt.Format("02-01-2006 15:04:05"),
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

func (h *TravelCostManagementHandler) GetTravelCostByID(c *gin.Context) {
	id := c.Param("id")
	travelCost, err := h.usecase.GetTravelCostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "travel cost not found"})
		return
	}

	// Map to TravelCostResponse DTO for clean response
	response := dto.TravelCostResponse{
		ID:        travelCost.ID,
		Name:      travelCost.Name,
		Code:      travelCost.Code,
		Unit:      travelCost.Unit,
		Price:     travelCost.Price,
		CreatedAt: travelCost.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: travelCost.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *TravelCostManagementHandler) GetEditTravelCostPage(c *gin.Context) {
	id := c.Param("id")

	// Fetch travel cost by ID
	travelCost, err := h.usecase.GetTravelCostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "travel cost not found"})
		return
	}

	// Map to TravelCostResponse DTO
	response := dto.TravelCostResponse{
		ID:        travelCost.ID,
		Name:      travelCost.Name,
		Code:      travelCost.Code,
		Unit:      travelCost.Unit,
		Price:     travelCost.Price,
		CreatedAt: travelCost.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: travelCost.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	// For travel costs, we don't have any reference data needed for editing
	// but we maintain the same response structure for consistency
	c.JSON(http.StatusOK, gin.H{
		"data": response,
		"reference": gin.H{
			// You can add any reference data here if needed in the future
			"units": []string{"km", "day", "night", "trip", "person"},
		},
	})
}

func (h *TravelCostManagementHandler) UpdateTravelCost(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	var input dto.UpdateTravelCostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travelCost := entity.TravelCost{
		ID:    uint(idUint),
		Name:  input.Name,
		Code:  input.Code,
		Unit:  input.Unit,
		Price: input.Price,
	}

	if err := h.usecase.UpdateTravelCost(id, &travelCost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "travel cost updated successfully", "data": travelCost})
}

func (h *TravelCostManagementHandler) DeleteTravelCost(c *gin.Context) {
	id := c.Param("id")

	if err := h.usecase.DeleteTravelCost(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "travel cost deleted successfully", "data": gin.H{"id": id}})
}

// GetAddTravelCostPage returns any reference data needed for the add travel cost page
func (h *TravelCostManagementHandler) GetAddTravelCostPage(c *gin.Context) {
	// For travel costs, we might want to provide a list of common units
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{},
		"reference": gin.H{
			"units": []string{"km", "day", "night", "trip", "person"},
		},
	})
}
