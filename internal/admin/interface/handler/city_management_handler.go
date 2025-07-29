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

type CityManagementHandler struct {
	usecase adminUsecase.CityManagementUsecase
}

func NewCityManagementHandler(r *gin.Engine, uc adminUsecase.CityManagementUsecase, db *gorm.DB) {
	h := &CityManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.GET("/cities/add", h.GetAddCityPage)
	admin.POST("/cities", h.CreateCity)
	admin.GET("/cities", h.GetCities)
	admin.GET("/cities/:id", h.GetCityByID)
	admin.GET("/cities/:id/edit", h.GetEditCityPage)
	admin.PUT("/cities/:id", h.UpdateCity)
	admin.DELETE("/cities/:id", h.DeleteCity)
}

func (h *CityManagementHandler) CreateCity(c *gin.Context) {
	var input dto.CreateCityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city := entity.City{
		Name:       input.Name,
		ProvinceID: &input.ProvinceID,
	}

	if err := h.usecase.CreateCity(&city); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "city created successfully", "data": city})
}

func (h *CityManagementHandler) GetAddCityPage(c *gin.Context) {
	// Fetch list of provinces
	provinces, err := h.usecase.GetProvinces(1, 100, "") // Assuming a method exists to fetch provinces
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var provinceList []dto.ProvinceResponse
	for _, province := range provinces {
		provinceList = append(provinceList, dto.ProvinceResponse{
			ID:   province.ID,
			Name: province.Name,
			CreatedAt: province.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: province.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	response := gin.H{
		"data": gin.H{
			"provinces": provinceList,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *CityManagementHandler) GetCities(c *gin.Context) {
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

	search := c.DefaultQuery("search", "") // Added search query parameter

	cities, err := h.usecase.GetCities(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get total count for pagination
	totalData, err := h.usecase.GetCitiesCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(totalData) / limit
	if int(totalData)%limit > 0 {
		totalPages++
	}

	var responseData []dto.CityResponse
	for _, city := range cities {
		responseData = append(responseData, dto.CityResponse{
			ID:       city.ID,
			Name:     city.Name,
			Province: city.Province.Name,
			CreatedAt: city.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: city.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
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

func (h *CityManagementHandler) GetCityByID(c *gin.Context) {
	id := c.Param("id")
	city, err := h.usecase.GetCityByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": city})
}

func (h *CityManagementHandler) UpdateCity(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid city ID"})
		return
	}

	var input dto.UpdateCityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city := entity.City{
		ID:         uint(idUint),
		Name:       input.Name,
		ProvinceID: &input.ProvinceID,
	}

	if err := h.usecase.UpdateCity(id, &city); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "city updated successfully", "data": city})
}

func (h *CityManagementHandler) DeleteCity(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteCity(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "city deleted successfully", "data": gin.H{"id": id}})
}

func (h *CityManagementHandler) GetEditCityPage(c *gin.Context) {
	id := c.Param("id")

	// Fetch city by ID
	city, err := h.usecase.GetCityByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch list of provinces
	provinces, err := h.usecase.GetProvinces(1, 100, "") // Assuming a method exists to fetch provinces
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var provinceList []dto.ProvinceResponse
	for _, province := range provinces {
		provinceList = append(provinceList, dto.ProvinceResponse{
			ID:   province.ID,
			Name: province.Name,
		})
	}

	response := gin.H{
		"data": city,
		"refrences": gin.H{
			"provinces": provinceList,
		},
	}

	c.JSON(http.StatusOK, response)
}
