package adminHandler

import (
	"net/http"
	"strconv"

	//    "fmt"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/middleware"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SupplierManagementHandler struct {
	usecase adminUsecase.SupplierManagementUsecase
}

func NewSupplierManagementHandler(r *gin.Engine, uc adminUsecase.SupplierManagementUsecase, db *gorm.DB) {
	h := &SupplierManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/suppliers", h.CreateSupplier)
	admin.GET("/suppliers", h.GetSuppliers)
	admin.GET("/suppliers/:id", h.GetSupplierByID)
	admin.PUT("/suppliers/:id", h.UpdateSupplier)
	admin.DELETE("/suppliers/:id", h.DeleteSupplier)
	admin.GET("/suppliers/:id/edit", h.GetSupplierEditPage)
}

func (h *SupplierManagementHandler) CreateSupplier(c *gin.Context) {
	var input dto.CreateSupplierInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier := entity.Supplier{
		UserID:        input.UserID,
		Name:          input.Name,
		InvoiceName:   input.InvoiceName,
		NPWP:          input.NPWP,
		Address:       input.Address,
		CityID:        input.CityID,
		Phone:         input.Phone,
		Email:         input.Email,
		Greeting:      input.Greeting,
		ContactPerson: input.ContactPerson,
		ContactPhone:  input.ContactPhone,
		BankAccount:   input.BankAccount,
		Type:          input.Type,
		LogoFilename:  input.LogoFilename,
	}

	if err := h.usecase.CreateSupplier(&supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "supplier created successfully", "data": supplier})
}

func (h *SupplierManagementHandler) GetSuppliers(c *gin.Context) {
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

	suppliers, err := h.usecase.GetSuppliers(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseData []dto.SupplierResponse
	for _, supplier := range suppliers {
		responseData = append(responseData, dto.SupplierResponse{
			ID:            supplier.ID,
			UserID:        supplier.UserID,
			Code:          "S" + strconv.FormatUint(uint64(supplier.Code), 10),
			Name:          supplier.Name,
			InvoiceName:   supplier.InvoiceName,
			NPWP:          supplier.NPWP,
			Address:       supplier.Address,
			City:          supplier.City.Name,
			Province:      supplier.City.Province.Name,
			Phone:         supplier.Phone,
			Email:         supplier.Email,
			Greeting:      supplier.Greeting,
			ContactPerson: supplier.ContactPerson,
			ContactPhone:  supplier.ContactPhone,
			BankAccount:   supplier.BankAccount,
			Type:          supplier.Type,
			LogoFilename:  supplier.LogoFilename,
			UpdatedAt:     supplier.UpdatedAt.Format("02-01-2006 15:04:05"), // Add if you have UpdatedAt field in entity.Supplier
		})
	}

	response := gin.H{
		"data": responseData,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *SupplierManagementHandler) GetSupplierByID(c *gin.Context) {
	id := c.Param("id")
	supplier, err := h.usecase.GetSupplierByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.SupplierResponse{
		ID:            supplier.ID,
		UserID:        supplier.UserID,
		Name:          supplier.Name,
		Code:          "S" + strconv.FormatUint(uint64(supplier.Code), 10),
		InvoiceName:   supplier.InvoiceName,
		NPWP:          supplier.NPWP,
		Address:       supplier.Address,
		City:          supplier.City.Name,
		Province:      supplier.City.Province.Name,
		Phone:         supplier.Phone,
		Email:         supplier.Email,
		Greeting:      supplier.Greeting,
		ContactPerson: supplier.ContactPerson,
		ContactPhone:  supplier.ContactPhone,
		BankAccount:   supplier.BankAccount,
		Type:          supplier.Type,
		LogoFilename:  supplier.LogoFilename,
		UpdatedAt:     "", // Add if you have UpdatedAt field in entity.Supplier
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *SupplierManagementHandler) UpdateSupplier(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid supplier ID"})
		return
	}

	var input dto.UpdateSupplierInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier := entity.Supplier{
		ID:            uint(idUint),
		UserID:        input.UserID,
		Name:          input.Name,
		InvoiceName:   input.InvoiceName,
		NPWP:          input.NPWP,
		Address:       input.Address,
		CityID:        input.CityID,
		Phone:         input.Phone,
		Email:         input.Email,
		Greeting:      input.Greeting,
		ContactPerson: input.ContactPerson,
		ContactPhone:  input.ContactPhone,
		BankAccount:   input.BankAccount,
		Type:          input.Type,
		LogoFilename:  input.LogoFilename,
	}

	if err := h.usecase.UpdateSupplier(id, &supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "supplier updated successfully", "data": supplier})
}

func (h *SupplierManagementHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteSupplier(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "supplier deleted successfully", "data": gin.H{"id": id}})
}

func (h *SupplierManagementHandler) GetSupplierEditPage(c *gin.Context) {
	id := c.Param("id")
	supplier, err := h.usecase.GetSupplierByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "supplier not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.SupplierResponse{
		ID:            supplier.ID,
		UserID:        supplier.UserID,
		Name:          supplier.Name,
		Code:          "S" + strconv.FormatUint(uint64(supplier.Code), 10),
		InvoiceName:   supplier.InvoiceName,
		NPWP:          supplier.NPWP,
		Address:       supplier.Address,
		City:          supplier.City.Name,
		Province:      supplier.City.Province.Name,
		Phone:         supplier.Phone,
		Email:         supplier.Email,
		Greeting:      supplier.Greeting,
		ContactPerson: supplier.ContactPerson,
		ContactPhone:  supplier.ContactPhone,
		BankAccount:   supplier.BankAccount,
		Type:          supplier.Type,
		LogoFilename:  supplier.LogoFilename,
		UpdatedAt:     supplier.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
