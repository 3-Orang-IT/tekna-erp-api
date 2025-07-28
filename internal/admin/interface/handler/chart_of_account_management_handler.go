package adminHandler

import (
	"net/http"
	"strconv"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChartOfAccountManagementHandler struct {
	usecase adminUsecase.ChartOfAccountManagementUsecase
}

func NewChartOfAccountManagementHandler(r *gin.Engine, uc adminUsecase.ChartOfAccountManagementUsecase, db *gorm.DB) {
	h := &ChartOfAccountManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.POST("/chart-of-accounts", h.CreateChartOfAccount)
	admin.GET("/chart-of-accounts", h.GetChartOfAccounts)
	admin.GET("/chart-of-accounts/:id", h.GetChartOfAccountByID)
	admin.GET("/chart-of-accounts/:id/edit", h.GetEditChartOfAccountPage)
	admin.PUT("/chart-of-accounts/:id", h.UpdateChartOfAccount)
	admin.DELETE("/chart-of-accounts/:id", h.DeleteChartOfAccount)
}

func (h *ChartOfAccountManagementHandler) CreateChartOfAccount(c *gin.Context) {
	var input dto.CreateChartOfAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chartOfAccount := entity.ChartOfAccount{
		Type: input.Type,
		Name: input.Name,
	}

	if err := h.usecase.CreateChartOfAccount(&chartOfAccount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Chart of Account created successfully", "data": chartOfAccount})
}

func (h *ChartOfAccountManagementHandler) GetChartOfAccounts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	search := c.DefaultQuery("search", "")

	chartOfAccounts, err := h.usecase.GetChartOfAccounts(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get total count for pagination
	totalData, err := h.usecase.GetChartOfAccountsCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(totalData) / limit
	if int(totalData)%limit > 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, gin.H{
		"data": chartOfAccounts, 
		"pagination": gin.H{
			"page": page, 
			"limit": limit,
			"total_data": totalData,
			"total_pages": totalPages,
		},
	})
}

func (h *ChartOfAccountManagementHandler) GetChartOfAccountByID(c *gin.Context) {
	id := c.Param("id")
	chartOfAccount, err := h.usecase.GetChartOfAccountByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chartOfAccount})
}

func (h *ChartOfAccountManagementHandler) UpdateChartOfAccount(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateChartOfAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chartOfAccount := entity.ChartOfAccount{
		Type: input.Type,
		Name: input.Name,
	}

	if err := h.usecase.UpdateChartOfAccount(id, &chartOfAccount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chart of Account updated successfully"})
}

func (h *ChartOfAccountManagementHandler) DeleteChartOfAccount(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteChartOfAccount(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chart of Account deleted successfully"})
}

func (h *ChartOfAccountManagementHandler) GetEditChartOfAccountPage(c *gin.Context) {
	id := c.Param("id")
	chartOfAccount, err := h.usecase.GetChartOfAccountByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chartOfAccount})
}
