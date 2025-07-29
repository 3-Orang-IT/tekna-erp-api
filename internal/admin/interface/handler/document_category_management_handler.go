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

type DocumentCategoryManagementHandler struct {
	usecase adminUsecase.DocumentCategoryManagementUsecase
}

func NewDocumentCategoryManagementHandler(r *gin.Engine, uc adminUsecase.DocumentCategoryManagementUsecase, db *gorm.DB) {
	h := &DocumentCategoryManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/document-categories", h.CreateDocumentCategory)
	admin.GET("/document-categories", h.GetDocumentCategories)
	admin.GET("/document-categories/add", h.GetAddDocumentCategoryPage)
	admin.GET("/document-categories/:id", h.GetDocumentCategoryByID)
	admin.GET("/document-categories/:id/edit", h.GetEditDocumentCategoryPage)
	admin.PUT("/document-categories/:id", h.UpdateDocumentCategory)
	admin.DELETE("/document-categories/:id", h.DeleteDocumentCategory)
}

func (h *DocumentCategoryManagementHandler) CreateDocumentCategory(c *gin.Context) {
	var input dto.CreateDocumentCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	documentCategory := entity.DocumentCategory{
		Name: input.Name,
	}

	if err := h.usecase.CreateDocumentCategory(&documentCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "document category created successfully", "data": documentCategory})
}

func (h *DocumentCategoryManagementHandler) GetDocumentCategories(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.DefaultQuery("search", "")

	// Get total count of document categories for pagination
	total, err := h.usecase.GetDocumentCategoriesCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	documentCategories, err := h.usecase.GetDocumentCategories(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map document categories to the desired response format using DocumentCategoryResponse DTO
	var responseData []dto.DocumentCategoryResponse
	for _, documentCategory := range documentCategories {
		responseData = append(responseData, dto.DocumentCategoryResponse{
			ID:        documentCategory.ID,
			Name:      documentCategory.Name,
			CreatedAt: documentCategory.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt: documentCategory.UpdatedAt.Format("02-01-2006 15:04:05"),
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

func (h *DocumentCategoryManagementHandler) GetDocumentCategoryByID(c *gin.Context) {
	id := c.Param("id")
	documentCategory, err := h.usecase.GetDocumentCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "document category not found"})
		return
	}

	// Map to DocumentCategoryResponse DTO for clean response
	response := dto.DocumentCategoryResponse{
		ID:        documentCategory.ID,
		Name:      documentCategory.Name,
		CreatedAt: documentCategory.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: documentCategory.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *DocumentCategoryManagementHandler) GetEditDocumentCategoryPage(c *gin.Context) {
	id := c.Param("id")

	// Fetch document category by ID
	documentCategory, err := h.usecase.GetDocumentCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "document category not found"})
		return
	}

	// Map to DocumentCategoryResponse DTO
	response := dto.DocumentCategoryResponse{
		ID:        documentCategory.ID,
		Name:      documentCategory.Name,
		CreatedAt: documentCategory.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: documentCategory.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	// For document categories, we don't have any reference data needed for editing
	// but we maintain the same response structure for consistency
	c.JSON(http.StatusOK, gin.H{
		"data": response,
		"reference": gin.H{}, // Empty since no reference data needed
	})
}

// GetAddDocumentCategoryPage returns any reference data needed for the add document category page
func (h *DocumentCategoryManagementHandler) GetAddDocumentCategoryPage(c *gin.Context) {
	// For document categories, we don't need any reference data
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{},
		"reference": gin.H{}, // Empty since no reference data needed
	})
}

func (h *DocumentCategoryManagementHandler) UpdateDocumentCategory(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	var input dto.UpdateDocumentCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	documentCategory := entity.DocumentCategory{
		ID:   uint(idUint),
		Name: input.Name,
	}

	if err := h.usecase.UpdateDocumentCategory(id, &documentCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "document category updated successfully", "data": documentCategory})
}

func (h *DocumentCategoryManagementHandler) DeleteDocumentCategory(c *gin.Context) {
	id := c.Param("id")

	if err := h.usecase.DeleteDocumentCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "document category deleted successfully", "data": gin.H{"id": id}})
}
