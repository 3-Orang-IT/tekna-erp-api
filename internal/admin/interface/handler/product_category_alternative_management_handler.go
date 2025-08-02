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

type ProductCategoryAlternativeManagementHandler struct {
	usecase adminUsecase.ProductCategoryAlternativeManagementUsecase
}

func NewProductCategoryAlternativeManagementHandler(r *gin.Engine, uc adminUsecase.ProductCategoryAlternativeManagementUsecase, db *gorm.DB) {
	h := &ProductCategoryAlternativeManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/product-category-alternatives", h.CreateProductCategoryAlternative)
	admin.GET("/product-category-alternatives", h.GetProductCategoryAlternatives)
	admin.GET("/product-category-alternatives/:id", h.GetProductCategoryAlternativeByID)
	admin.GET("/product-category-alternatives/:id/edit", h.GetEditProductCategoryAlternativePage)
	admin.PUT("/product-category-alternatives/:id", h.UpdateProductCategoryAlternative)
	admin.DELETE("/product-category-alternatives/:id", h.DeleteProductCategoryAlternative)
}

func (h *ProductCategoryAlternativeManagementHandler) CreateProductCategoryAlternative(c *gin.Context) {
	var input dto.CreateProductCategoryAlternativeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := entity.ProductCategoryAlternative{
		Name: input.Name,
	}
	if err := h.usecase.CreateProductCategoryAlternative(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "product category alternative created successfully", "data": category})
}

func (h *ProductCategoryAlternativeManagementHandler) GetProductCategoryAlternatives(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}
	search := c.DefaultQuery("search", "")
	total, err := h.usecase.GetProductCategoryAlternativesCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}
	categories, err := h.usecase.GetProductCategoryAlternatives(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var responseData []dto.ProductCategoryAlternativeResponse
	for _, category := range categories {
		responseData = append(responseData, dto.ProductCategoryAlternativeResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
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

func (h *ProductCategoryAlternativeManagementHandler) GetProductCategoryAlternativeByID(c *gin.Context) {
	id := c.Param("id")
	category, err := h.usecase.GetProductCategoryAlternativeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product category alternative not found"})
		return
	}
	response := dto.ProductCategoryAlternativeResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *ProductCategoryAlternativeManagementHandler) GetEditProductCategoryAlternativePage(c *gin.Context) {
	id := c.Param("id")
	category, err := h.usecase.GetProductCategoryAlternativeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product category alternative not found"})
		return
	}
	response := dto.ProductCategoryAlternativeResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, gin.H{"data": response, "reference": gin.H{}})
}

func (h *ProductCategoryAlternativeManagementHandler) UpdateProductCategoryAlternative(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateProductCategoryAlternativeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := entity.ProductCategoryAlternative{
		Name: input.Name,
	}
	if err := h.usecase.UpdateProductCategoryAlternative(id, &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product category alternative updated successfully", "data": category})
}

func (h *ProductCategoryAlternativeManagementHandler) DeleteProductCategoryAlternative(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteProductCategoryAlternative(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product category alternative deleted successfully", "data": gin.H{"id": id}})
}
