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

type ProvinceManagementHandler struct {
	usecase adminUsecase.ProvinceManagementUsecase
}

func NewProvinceManagementHandler(r *gin.Engine, uc adminUsecase.ProvinceManagementUsecase, db *gorm.DB) {
	   h := &ProvinceManagementHandler{uc}
	   admin := r.Group("/api/v1/admin")
	   admin.Use(middleware.AdminRoleMiddleware(db))
	   admin.POST("/provinces", h.CreateProvince)
	   admin.GET("/provinces", h.GetProvinces)
	   admin.GET("/provinces/:id", h.GetProvinceByID)
	   admin.GET("/provinces/:id/edit", h.GetProvinceEditPage)
	   admin.PUT("/provinces/:id", h.UpdateProvince)
	   admin.DELETE("/provinces/:id", h.DeleteProvince)
}

// GetProvinceEditPage returns province data for the edit page
func (h *ProvinceManagementHandler) GetProvinceEditPage(c *gin.Context) {
	   id := c.Param("id")
	   province, err := h.usecase.GetProvinceByID(id)
	   if err != nil {
			   c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			   return
	   }
	   c.JSON(http.StatusOK, gin.H{"data": province})
}


func (h *ProvinceManagementHandler) CreateProvince(c *gin.Context) {
	var input dto.CreateProvinceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	province := entity.Province{
		Name: input.Name,
	}

	if err := h.usecase.CreateProvince(&province); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "province created successfully", "data": province})
}

func (h *ProvinceManagementHandler) GetProvinces(c *gin.Context) {
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

	   provinces, err := h.usecase.GetProvinces(page, limit, search)
	   if err != nil {
			   c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			   return
	   }
	   var responseData []dto.ProvinceResponse
	   for _, province := range provinces {
		   responseData = append(responseData, dto.ProvinceResponse{
			   ID:   province.ID,
			   Name: province.Name,
			   CreatedAt: province.CreatedAt.Format("2006-01-02 15:04:05"),
			   UpdatedAt: province.UpdatedAt.Format("2006-01-02 15:04:05"),
		   })
	   }

	   // Get total count for pagination
	   totalData, err := h.usecase.GetProvincesCount(search)
	   if err != nil {
			   c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			   return
	   }

	   // Calculate total pages
	   totalPages := int(totalData) / limit
	   if int(totalData)%limit > 0 {
			   totalPages++
	   }

	   response := gin.H{
			   "data": responseData,
			   "pagination": gin.H{
					   "page":        page,
					   "limit":       limit,
					   "total_data":  totalData,
					   "total_pages": totalPages,
			   },
	   }

	   c.JSON(http.StatusOK, response)
}

func (h *ProvinceManagementHandler) GetProvinceByID(c *gin.Context) {
	id := c.Param("id")
	province, err := h.usecase.GetProvinceByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": province})
}

func (h *ProvinceManagementHandler) UpdateProvince(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid province ID"})
		return
	}

	var input dto.UpdateProvinceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	province := entity.Province{
		ID:   uint(idUint),
		Name: input.Name,
	}

	if err := h.usecase.UpdateProvince(id, &province); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "province updated successfully", "data": province})
}

func (h *ProvinceManagementHandler) DeleteProvince(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteProvince(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "province deleted successfully", "data": gin.H{"id": id}})
}
