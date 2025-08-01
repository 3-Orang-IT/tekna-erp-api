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

type ProductManagementHandler struct {
	usecase adminUsecase.ProductManagementUsecase
}

func NewProductManagementHandler(r *gin.Engine, uc adminUsecase.ProductManagementUsecase, db *gorm.DB) {
	h := &ProductManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/products", h.CreateProduct)
	admin.GET("/products", h.GetProducts)
	admin.GET("/products/add", h.GetAddProductPage) // New route for add page
	admin.GET("/products/:id", h.GetProductByID)
	admin.GET("/products/:id/edit", h.GetEditProductPage)
	admin.PUT("/products/:id", h.UpdateProduct)
	admin.DELETE("/products/:id", h.DeleteProduct)
}

func (h *ProductManagementHandler) CreateProduct(c *gin.Context) {
	var input dto.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := entity.Product{
		ProductCategoryID: &input.ProductCategoryID,
		SupplierID:        input.SupplierID,
		BusinessUnitID:    input.BusinessUnitID,
		UnitID:            input.UnitID,
		Barcode:           input.Barcode,
		Name:              input.Name,
		Description:       input.Description,
		MaxQuantity:       input.MaxQuantity,
		MinQuantity:       input.MinQuantity,
		PurchasePrice:     input.PurchasePrice,
		SellingPrice:      input.SellingPrice,
		IsRecommended:     input.IsRecommended,
		ProductFocus:      input.ProductFocus,
		Brand:             input.Brand,
	}

	// Generate an auto-incrementing product code
	lastProduct, err := h.usecase.GetLastProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate product code"})
		return
	}
	product.Code = "PRD-" + strconv.Itoa(int(lastProduct.ID)+1)

	if err := h.usecase.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.ProductResponse{
		ID:                product.ID,
		Code:              product.Code,
		Barcode:           product.Barcode,
		Name:              product.Name,
		ProductCategoryName: product.ProductCategory.Name,
		SupplierName:      product.Supplier.Name,
		BusinessUnitName:  product.BusinessUnit.Name,
		UnitName:          product.Unit.Name,
		Description:       product.Description,
		MaxQuantity:       product.MaxQuantity,
		MinQuantity:       product.MinQuantity,
		PurchasePrice:     product.PurchasePrice,
		SellingPrice:      product.SellingPrice,
		IsRecommended:     product.IsRecommended,
		ProductFocus:      product.ProductFocus,
		Brand:             product.Brand,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "product created successfully", "data": response})
}

func (h *ProductManagementHandler) GetProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number"})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit number"})
		return
	}
	search := c.DefaultQuery("search", "")
	
	// Get total count of products for pagination
	total, err := h.usecase.GetProductsCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}
	
	products, err := h.usecase.GetProducts(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Map products to the desired response format using ProductResponse DTO
	var responseData []dto.ProductResponse
	for _, product := range products {
		responseData = append(responseData, dto.ProductResponse{
			ID:                product.ID,
			Code:              product.Code,
			Barcode:           product.Barcode,
			Name:              product.Name,
			ProductCategoryName: product.ProductCategory.Name,
			SupplierName:      product.Supplier.Name,
			BusinessUnitName:  product.BusinessUnit.Name,
			UnitName:          product.Unit.Name,
			Description:       product.Description,
			MaxQuantity:       product.MaxQuantity,
			MinQuantity:       product.MinQuantity,
			PurchasePrice:     product.PurchasePrice,
			SellingPrice:      product.SellingPrice,
			IsRecommended:     product.IsRecommended,
			ProductFocus:      product.ProductFocus,
			Brand:             product.Brand,
			CreatedAt:         product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:         product.UpdatedAt.Format("2006-01-02 15:04:05"),	
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

func (h *ProductManagementHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := h.usecase.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (h *ProductManagementHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var input entity.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.UpdateProduct(id, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully", "data": input})
}

func (h *ProductManagementHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}

func (h *ProductManagementHandler) GetEditProductPage(c *gin.Context) {
	id := c.Param("id")
	product, err := h.usecase.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Fetch additional data for the edit page
	productCategories, err := h.usecase.GetProductCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product categories"})
		return
	}
	suppliers, err := h.usecase.GetSuppliers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch suppliers"})
		return
	}

	businessUnits, err := h.usecase.GetBusinessUnits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch business units"})
		return
	}

	units, err := h.usecase.GetUnits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch units"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
		"references": gin.H{
			"product_categories": productCategories,
			"suppliers": suppliers,
			"business_units": businessUnits,
			"units": units,
		},
	})
}

// GetAddProductPage returns necessary reference data for the add product page
func (h *ProductManagementHandler) GetAddProductPage(c *gin.Context) {
	// Fetch product categories
	productCategories, err := h.usecase.GetProductCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product categories"})
		return
	}

	// Fetch suppliers
	suppliers, err := h.usecase.GetSuppliers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch suppliers"})
		return
	}

	// Fetch business units
	businessUnits, err := h.usecase.GetBusinessUnits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch business units"})
		return
	}

	// Fetch units of measure
	units, err := h.usecase.GetUnits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch units"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"product_categories": productCategories,
			"suppliers": suppliers,
			"business_units": businessUnits,
			"units": units,
		},
	})
}
