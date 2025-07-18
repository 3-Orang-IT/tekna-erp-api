package handler

import (
    "errors"
    "net/http"
    "strings"

    "github.com/3-Orang-IT/tekna-erp-api/internal/admin/middleware"
    "github.com/3-Orang-IT/tekna-erp-api/internal/company/usecase"
    "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type CompanyHandler struct {
    usecase usecase.CompanyUsecase
}

func NewCompanyHandler(r *gin.Engine, uc usecase.CompanyUsecase, db *gorm.DB) {
    h := &CompanyHandler{uc}
    
    // Public routes (untuk get company info)
    api := r.Group("/api/v1")
    api.GET("/company", h.GetCompany)
    
    // Protected routes (untuk admin operations)
    admin := r.Group("/api/v1/admin")
    admin.Use(middleware.AdminRoleMiddleware(db))
    admin.POST("/company", h.CreateCompany)
    admin.PUT("/company/:id", h.UpdateCompany)
    admin.DELETE("/company/:id", h.DeleteCompany)
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
    company, err := h.usecase.GetCompany()
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Company retrieved successfully",
        "data":    company,
    })
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
    var company entity.Company
    if err := c.ShouldBindJSON(&company); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validateCompany(&company); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.usecase.CreateCompany(&company); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Company created successfully",
        "data":    company,
    })
}

func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
    id := c.Param("id")
    
    // Get existing company first
    existingCompany, err := h.usecase.GetCompany()
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Check if the ID matches
    if existingCompany.ID != id {
        c.JSON(http.StatusNotFound, gin.H{"error": "Company ID not found"})
        return
    }

    var company entity.Company
    if err := c.ShouldBindJSON(&company); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Set the ID to ensure we're updating the correct record
    company.ID = id

    if err := h.usecase.UpdateCompany(&company); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Company updated successfully",
        "data":    company,
    })
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
    id := c.Param("id")
    
    if err := h.usecase.DeleteCompany(id); err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Company ID not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Company deleted successfully",
        "data":    gin.H{"id": id},
    })
}

func validateCompany(company *entity.Company) error {
    if strings.TrimSpace(company.Name) == "" {
        return errors.New("company name is required")
    }
    if strings.TrimSpace(company.Address) == "" {
        return errors.New("company address is required")
    }
    if strings.TrimSpace(company.Email) == "" {
        return errors.New("company email is required")
    }
    return nil
}