package adminHandler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/middleware"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewsletterManagementHandler struct {
	usecase adminUsecase.NewsletterManagementUsecase
}

func NewNewsletterManagementHandler(r *gin.Engine, uc adminUsecase.NewsletterManagementUsecase, db *gorm.DB) {
	h := &NewsletterManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/newsletters", h.CreateNewsletter)
	admin.GET("/newsletters", h.GetNewsletters)
	admin.GET("/newsletters/add", h.GetAddNewsletterPage)
	admin.GET("/newsletters/:id", h.GetNewsletterByID)
	admin.GET("/newsletters/:id/edit", h.GetEditNewsletterPage)
	admin.PUT("/newsletters/:id", h.UpdateNewsletter)
	admin.DELETE("/newsletters/:id", h.DeleteNewsletter)
}

// ensureDirNewsletter creates a directory if it does not exist
func ensureDirNewsletter(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func (h *NewsletterManagementHandler) CreateNewsletter(c *gin.Context) {
	// Parse form-data (multipart)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form: " + err.Error()})
		return
	}

	// Get fields from form-data
	newsletterType := c.PostForm("type")
	title := c.PostForm("title")
	description := c.PostForm("description")
	validFrom := c.PostForm("valid_from")
	status := c.PostForm("status")

	// Handle file upload
	file, header, err := c.Request.FormFile("file")
	var filePath string
	if err == nil && header != nil {
		// Sanitize filename to prevent directory traversal or invalid chars
		sanitizedFilename := strings.ReplaceAll(header.Filename, " ", "_")
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), sanitizedFilename)
		savePath := fmt.Sprintf("uploads/newsletters/%s", filename)
		
		// Ensure directory exists
		if err := ensureDirNewsletter("uploads/newsletters"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upload directory: " + err.Error()})
			return
		}
		
		out, err := os.Create(savePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file: " + err.Error()})
			return
		}
		defer out.Close()
		
		if _, err := io.Copy(out, file); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write file: " + err.Error()})
			return
		}
		
		filePath = savePath
	}

	newsletter := entity.Newsletter{
		Type:        newsletterType,
		Title:       title,
		Description: description,
		File:        filePath,
		ValidFrom:   validFrom,
		Status:      status,
	}

	if err := h.usecase.CreateNewsletter(&newsletter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "newsletter created successfully", "data": newsletter})
}

func (h *NewsletterManagementHandler) GetNewsletters(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.DefaultQuery("search", "")

	// Get total count of newsletters for pagination
	total, err := h.usecase.GetNewslettersCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	newsletters, err := h.usecase.GetNewsletters(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map newsletters to the desired response format using NewsletterResponse DTO
	var responseData []dto.NewsletterResponse
	baseUrl := os.Getenv("BASE_URL")
	for _, newsletter := range newsletters {
		var fileURL string
		if newsletter.File != "" {
			fileURL = baseUrl + newsletter.File
		} else {
			fileURL = ""
		}
		
		responseData = append(responseData, dto.NewsletterResponse{
			ID:          newsletter.ID,
			Type:        newsletter.Type,
			Title:       newsletter.Title,
			Description: newsletter.Description,
			File:        fileURL,
			ValidFrom:   newsletter.ValidFrom,
			Status:      newsletter.Status,
			CreatedAt:   newsletter.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt:   newsletter.UpdatedAt.Format("02-01-2006 15:04:05"),
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

func (h *NewsletterManagementHandler) GetNewsletterByID(c *gin.Context) {
	id := c.Param("id")
	newsletter, err := h.usecase.GetNewsletterByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "newsletter not found"})
		return
	}

	// Format file URL with BASE_URL if available
	baseUrl := os.Getenv("BASE_URL")
	var fileURL string
	if newsletter.File != "" {
		fileURL = baseUrl + newsletter.File
	} else {
		fileURL = ""
	}

	// Map to NewsletterResponse DTO for clean response
	response := dto.NewsletterResponse{
		ID:          newsletter.ID,
		Type:        newsletter.Type,
		Title:       newsletter.Title,
		Description: newsletter.Description,
		File:        fileURL,
		ValidFrom:   newsletter.ValidFrom,
		Status:      newsletter.Status,
		CreatedAt:   newsletter.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:   newsletter.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *NewsletterManagementHandler) GetEditNewsletterPage(c *gin.Context) {
	id := c.Param("id")

	// Fetch newsletter by ID
	newsletter, err := h.usecase.GetNewsletterByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "newsletter not found"})
		return
	}

	// Format file URL with BASE_URL if available
	baseUrl := os.Getenv("BASE_URL")
	var fileURL string
	if newsletter.File != "" {
		fileURL = baseUrl + newsletter.File
	} else {
		fileURL = ""
	}

	response := dto.NewsletterResponse{
		ID:          newsletter.ID,
		Type:        newsletter.Type,
		Title:       newsletter.Title,
		Description: newsletter.Description,
		File:        fileURL,
		ValidFrom:   newsletter.ValidFrom,
		Status:      newsletter.Status,
		CreatedAt:   newsletter.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:   newsletter.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	// Include reference data for the edit page
	c.JSON(http.StatusOK, gin.H{
		"data": response,
		"reference": gin.H{
			"types": []string{"Announcement", "Promo", "News", "Update"},
			"statuses": []string{"Active", "Draft", "Archived"},
		},
	})
}

// GetAddNewsletterPage returns any reference data needed for the add newsletter page
func (h *NewsletterManagementHandler) GetAddNewsletterPage(c *gin.Context) {
	// For newsletters, we might not need any reference data
	// But if needed, you could return predefined types, statuses, etc.
	c.JSON(http.StatusOK, gin.H{
		"types": []string{"Announcement", "Promo", "News", "Update"},
		"statuses": []string{"Active", "Draft", "Archived"},
	})
}

func (h *NewsletterManagementHandler) UpdateNewsletter(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	// Parse form-data (multipart)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form: " + err.Error()})
		return
	}

	// Get fields from form-data
	newsletterType := c.PostForm("type")
	title := c.PostForm("title")
	description := c.PostForm("description")
	validFrom := c.PostForm("valid_from")
	status := c.PostForm("status")

	// Get existing newsletter to check if we need to keep the old file
	existingNewsletter, err := h.usecase.GetNewsletterByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "newsletter not found"})
		return
	}

	// Handle file upload
	file, header, err := c.Request.FormFile("file")
	var filePath string
	if err == nil && header != nil {
		// New file uploaded, replace the old one
		sanitizedFilename := strings.ReplaceAll(header.Filename, " ", "_")
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), sanitizedFilename)
		savePath := fmt.Sprintf("uploads/newsletters/%s", filename)
		
		// Ensure directory exists
		if err := ensureDirNewsletter("uploads/newsletters"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upload directory: " + err.Error()})
			return
		}
		
		out, err := os.Create(savePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file: " + err.Error()})
			return
		}
		defer out.Close()
		
		if _, err := io.Copy(out, file); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write file: " + err.Error()})
			return
		}
		
		// Delete old file if it exists
		if existingNewsletter.File != "" {
			_ = os.Remove(existingNewsletter.File) // Best effort deletion, ignore errors
		}
		
		filePath = savePath
	} else {
		// If no new file uploaded, keep the old file path
		filePath = existingNewsletter.File
	}

	newsletter := entity.Newsletter{
		ID:          uint(idUint),
		Type:        newsletterType,
		Title:       title,
		Description: description,
		File:        filePath,
		ValidFrom:   validFrom,
		Status:      status,
	}

	if err := h.usecase.UpdateNewsletter(id, &newsletter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "newsletter updated successfully", "data": newsletter})
}

func (h *NewsletterManagementHandler) DeleteNewsletter(c *gin.Context) {
	id := c.Param("id")

	// Get newsletter to access file path before deletion
	newsletter, err := h.usecase.GetNewsletterByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "newsletter not found"})
		return
	}

	// Delete file if it exists
	if newsletter.File != "" {
		_ = os.Remove(newsletter.File) // Best effort deletion, ignore errors
	}

	// Delete record from database
	if err := h.usecase.DeleteNewsletter(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "newsletter deleted successfully", "data": gin.H{"id": id}})
}
