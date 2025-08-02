package adminHandler

import (
	"net/http"
	"strconv"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
)

type ProductCategoryManagementHandler struct {
	usecase adminUsecase.ProductCategoryManagementUsecase
}

func NewProductCategoryManagementHandler(r *gin.Engine, uc adminUsecase.ProductCategoryManagementUsecase) {
	h := &ProductCategoryManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.POST("/product-categories", h.CreateProductCategory)
	admin.GET("/product-categories", h.GetProductCategories)
	admin.GET("/product-categories/:id", h.GetProductCategoryByID)
	admin.GET("/product-categories/:id/edit", h.GetProductCategoryByID)
	admin.PUT("/product-categories/:id", h.UpdateProductCategory)
	admin.DELETE("/product-categories/:id", h.DeleteProductCategory)
}

func (h *ProductCategoryManagementHandler) CreateProductCategory(c *gin.Context) {
	var input dto.CreateProductCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := entity.ProductCategory{
		Name: input.Name,
	}
	if err := h.usecase.CreateProductCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "product category created successfully", "data": category})
}

func (h *ProductCategoryManagementHandler) GetProductCategories(c *gin.Context) {
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
	
	// Get total count of product categories for pagination
	total, err := h.usecase.GetProductCategoriesCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}
	
	categories, err := h.usecase.GetProductCategories(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var responseData []dto.ProductCategoryResponse
	for _, category := range categories {
		responseData = append(responseData, dto.ProductCategoryResponse{
			ID: category.ID,
			Name: category.Name,
			CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
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

func (h *ProductCategoryManagementHandler) GetProductCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := h.usecase.GetProductCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := dto.ProductCategoryResponse{
		ID: category.ID,
		Name: category.Name,
		CreatedAt: category.CreatedAt.Format("02-01-2006 15:04"),
		UpdatedAt: category.UpdatedAt.Format("02-01-2006 15:04"),
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *ProductCategoryManagementHandler) UpdateProductCategory(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product category ID"})
		return
	}
	var input dto.UpdateProductCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := entity.ProductCategory{
		ID:   uint(idUint),
		Name: input.Name,
	}
	if err := h.usecase.UpdateProductCategory(id, &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product category updated successfully", "data": category})
}

func (h *ProductCategoryManagementHandler) DeleteProductCategory(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteProductCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product category deleted successfully", "data": gin.H{"id": id}})
}
