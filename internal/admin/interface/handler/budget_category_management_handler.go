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

type BudgetCategoryManagementHandler struct {
	usecase adminUsecase.BudgetCategoryManagementUsecase
}

func NewBudgetCategoryManagementHandler(r *gin.Engine, uc adminUsecase.BudgetCategoryManagementUsecase, db *gorm.DB) {
	h := &BudgetCategoryManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/budget-categories", h.CreateBudgetCategory)
	admin.GET("/budget-categories", h.GetBudgetCategories)
	admin.GET("/budget-categories/add", h.GetAddBudgetCategoryPage)
	admin.GET("/budget-categories/:id", h.GetBudgetCategoryByID)
	admin.GET("/budget-categories/:id/edit", h.GetEditBudgetCategoryPage)
	admin.PUT("/budget-categories/:id", h.UpdateBudgetCategory)
	admin.DELETE("/budget-categories/:id", h.DeleteBudgetCategory)
}

// GetAddBudgetCategoryPage returns reference data for the add budget category page

func (h *BudgetCategoryManagementHandler) CreateBudgetCategory(c *gin.Context) {
	var input dto.CreateBudgetCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := entity.BudgetCategory{
		ChartOfAccountID: input.ChartOfAccountID,
		Name:             input.Name,
		Description:      input.Description,
		Order:            input.Order,
	}

	if err := h.usecase.CreateBudgetCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "budget category created successfully", "data": category})
}

func (h *BudgetCategoryManagementHandler) GetBudgetCategories(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}
	search := c.DefaultQuery("search", "")

	total, err := h.usecase.GetBudgetCategoriesCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	categories, err := h.usecase.GetBudgetCategories(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseData []dto.BudgetCategoryResponse
	for _, category := range categories {
		responseData = append(responseData, dto.BudgetCategoryResponse{
			ID:               category.ID,
			ChartOfAccount:   getChartOfAccountName(category.ChartOfAccount),
			Name:             category.Name,
			Description:      category.Description,
			Order:            category.Order,
			CreatedAt:        category.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        category.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responseData,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total_data":  total,
			"total_pages": totalPages,
		},
	})
}

func getChartOfAccountName(account *entity.ChartOfAccount) string {
	if account != nil {
		return account.Name
	}
	return ""
}

func (h *BudgetCategoryManagementHandler) GetBudgetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := h.usecase.GetBudgetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "budget category not found"})
		return
	}
	response := dto.BudgetCategoryResponse{
		ID:               category.ID,
		ChartOfAccount:   getChartOfAccountName(category.ChartOfAccount),
		Name:             category.Name,
		Description:      category.Description,
		Order:            category.Order,
		CreatedAt:        category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        category.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *BudgetCategoryManagementHandler) GetEditBudgetCategoryPage(c *gin.Context) {
	id := c.Param("id")
	category, err := h.usecase.GetBudgetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "budget category not found"})
		return
	}
	accounts, err := h.usecase.GetChartOfAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch chart of accounts"})
		return
	}
	var accountRefs []gin.H
	for _, acc := range accounts {
		accountRefs = append(accountRefs, gin.H{"id": acc.ID, "name": acc.Name, "code": acc.Code})
	}
	response := dto.BudgetCategoryResponse{
		ID:               category.ID,
		ChartOfAccount:   getChartOfAccountName(category.ChartOfAccount),
		Name:             category.Name,
		Description:      category.Description,
		Order:            category.Order,
		CreatedAt:        category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        category.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, gin.H{"data": response, "reference": gin.H{"chart_of_accounts": accountRefs}})
}

func (h *BudgetCategoryManagementHandler) UpdateBudgetCategory(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateBudgetCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := entity.BudgetCategory{
		ChartOfAccountID: input.ChartOfAccountID,
		Name:             input.Name,
		Description:      input.Description,
		Order:            input.Order,
	}
	if err := h.usecase.UpdateBudgetCategory(id, &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "budget category updated successfully", "data": category})
}

func (h *BudgetCategoryManagementHandler) DeleteBudgetCategory(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteBudgetCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "budget category deleted successfully", "data": gin.H{"id": id}})
}

// GetAddBudgetCategoryPage returns reference data for the add budget category page
func (h *BudgetCategoryManagementHandler) GetAddBudgetCategoryPage(c *gin.Context) {
	accounts, err := h.usecase.GetChartOfAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch chart of accounts"})
		return
	}
	var accountRefs []gin.H
	for _, acc := range accounts {
		accountRefs = append(accountRefs, gin.H{"id": acc.ID, "name": acc.Name, "code": acc.Code})
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"chart_of_accounts": accountRefs}})
}