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

type DocumentManagementHandler struct {
	usecase adminUsecase.DocumentManagementUsecase
}

func NewDocumentManagementHandler(r *gin.Engine, uc adminUsecase.DocumentManagementUsecase, db *gorm.DB) {
	h := &DocumentManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/documents", h.CreateDocument)
	admin.GET("/documents", h.GetDocuments)
	admin.GET("/documents/add", h.GetAddDocumentPage)
	admin.GET("/documents/:id", h.GetDocumentByID)
	admin.GET("/documents/:id/edit", h.GetEditDocumentPage)
	admin.PUT("/documents/:id", h.UpdateDocument)
	admin.DELETE("/documents/:id", h.DeleteDocument)
}

// ensureDirDocument creates a directory if it does not exist
func ensureDirDocument(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func (h *DocumentManagementHandler) CreateDocument(c *gin.Context) {
	// Get the current user ID from the context (set by the middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Parse form-data (multipart)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form: " + err.Error()})
		return
	}

	// Get form fields
	documentCategoryIDStr := c.PostForm("document_category_id")
	documentCategoryID, err := strconv.ParseUint(documentCategoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid document category ID"})
		return
	}

	name := c.PostForm("name")
	description := c.PostForm("description")
	isPublishedStr := c.PostForm("is_published")
	isPublished := isPublishedStr == "true"

	// Handle file upload
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required: " + err.Error()})
		return
	}

	// Save the file
	sanitizedFilename := strings.ReplaceAll(header.Filename, " ", "_")
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), sanitizedFilename)
	savePath := fmt.Sprintf("uploads/documents/%s", filename)
	
	// Ensure directory exists
	if err := ensureDirDocument("uploads/documents"); err != nil {
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

	document := entity.Document{
		DocumentCategoryID: uint(documentCategoryID),
		Name:               name,
		UserID:             userID.(uint),
		FilePath:           savePath,
		Description:        description,
		IsPublished:        isPublished,
	}

	if err := h.usecase.CreateDocument(&document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "document created successfully", "data": document})
}

func (h *DocumentManagementHandler) GetDocuments(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.DefaultQuery("search", "")

	// Get total count of documents for pagination
	total, err := h.usecase.GetDocumentsCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	documents, err := h.usecase.GetDocuments(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map documents to the desired response format using DocumentResponse DTO
	var responseData []dto.DocumentResponse
	baseUrl := os.Getenv("BASE_URL")
	for _, document := range documents {
		var fileURL string
		if document.FilePath != "" {
			fileURL = baseUrl + document.FilePath
		}
		
		responseData = append(responseData, dto.DocumentResponse{
			ID:                 document.ID,
			DocumentCategoryID: document.DocumentCategoryID,
			DocumentCategory: dto.DocumentCategoryResponseLimited{
				ID:   document.DocumentCategory.ID,
				Name: document.DocumentCategory.Name,
			},
			Name:        document.Name,
			UserID:      document.UserID,
			User: dto.UserResponseLimited{
				ID:       document.User.ID,
				Username: document.User.Username,
				Name:     document.User.Name,
			},
			FilePath:    fileURL,
			Description: document.Description,
			IsPublished: document.IsPublished,
			CreatedAt:   document.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt:   document.UpdatedAt.Format("02-01-2006 15:04:05"),
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

func (h *DocumentManagementHandler) GetDocumentByID(c *gin.Context) {
	id := c.Param("id")
	document, err := h.usecase.GetDocumentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
		return
	}

	// Format file URL with BASE_URL if available
	baseUrl := os.Getenv("BASE_URL")
	var fileURL string
	if document.FilePath != "" {
		fileURL = baseUrl + document.FilePath
	}

	// Map to DocumentResponse DTO for clean response
	response := dto.DocumentResponse{
		ID:                 document.ID,
		DocumentCategoryID: document.DocumentCategoryID,
		DocumentCategory: dto.DocumentCategoryResponseLimited{
			ID:   document.DocumentCategory.ID,
			Name: document.DocumentCategory.Name,
		},
		Name:        document.Name,
		UserID:      document.UserID,
		User: dto.UserResponseLimited{
			ID:       document.User.ID,
			Username: document.User.Username,
			Name:     document.User.Name,
		},
		FilePath:    fileURL,
		Description: document.Description,
		IsPublished: document.IsPublished,
		CreatedAt:   document.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:   document.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *DocumentManagementHandler) GetEditDocumentPage(c *gin.Context) {
	id := c.Param("id")

	// Fetch document by ID
	document, err := h.usecase.GetDocumentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
		return
	}

	// Format file URL with BASE_URL if available
	baseUrl := os.Getenv("BASE_URL")
	var fileURL string
	if document.FilePath != "" {
		fileURL = baseUrl + document.FilePath
	}

	// Fetch document categories for reference
	documentCategories, err := h.usecase.GetDocumentCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch document categories"})
		return
	}

	// Map document categories to response format
	var categoryResponses []dto.DocumentCategoryResponseLimited
	for _, category := range documentCategories {
		categoryResponses = append(categoryResponses, dto.DocumentCategoryResponseLimited{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	// Map to DocumentResponse DTO
	documentResponse := dto.DocumentResponse{
		ID:                 document.ID,
		DocumentCategoryID: document.DocumentCategoryID,
		DocumentCategory: dto.DocumentCategoryResponseLimited{
			ID:   document.DocumentCategory.ID,
			Name: document.DocumentCategory.Name,
		},
		Name:        document.Name,
		UserID:      document.UserID,
		User: dto.UserResponseLimited{
			ID:       document.User.ID,
			Username: document.User.Username,
			Name:     document.User.Name,
		},
		FilePath:    fileURL,
		Description: document.Description,
		IsPublished: document.IsPublished,
		CreatedAt:   document.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:   document.UpdatedAt.Format("02-01-2006 15:04:05"),
	}

	// Return document data with reference data
	c.JSON(http.StatusOK, gin.H{
		"data": documentResponse,
		"reference": gin.H{
			"document_categories": categoryResponses,
		},
	})
}

// GetAddDocumentPage returns reference data needed for the add document page
func (h *DocumentManagementHandler) GetAddDocumentPage(c *gin.Context) {
	// Fetch document categories for reference
	documentCategories, err := h.usecase.GetDocumentCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch document categories"})
		return
	}

	// Map document categories to response format
	var categoryResponses []dto.DocumentCategoryResponseLimited
	for _, category := range documentCategories {
		categoryResponses = append(categoryResponses, dto.DocumentCategoryResponseLimited{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"document_categories": categoryResponses,
		},
	})
}

func (h *DocumentManagementHandler) UpdateDocument(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	// Get existing document to check if we need to keep the old file
	existingDocument, err := h.usecase.GetDocumentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
		return
	}

	// Parse form-data (multipart)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form: " + err.Error()})
		return
	}

	// Get form fields
	documentCategoryIDStr := c.PostForm("document_category_id")
	documentCategoryID, err := strconv.ParseUint(documentCategoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid document category ID"})
		return
	}

	name := c.PostForm("name")
	description := c.PostForm("description")
	isPublishedStr := c.PostForm("is_published")
	isPublished := isPublishedStr == "true"

	// Handle file upload
	file, header, err := c.Request.FormFile("file")
	var filePath string
	if err == nil && header != nil {
		// New file uploaded, replace the old one
		sanitizedFilename := strings.ReplaceAll(header.Filename, " ", "_")
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), sanitizedFilename)
		savePath := fmt.Sprintf("uploads/documents/%s", filename)
		
		// Ensure directory exists
		if err := ensureDirDocument("uploads/documents"); err != nil {
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
		if existingDocument.FilePath != "" {
			_ = os.Remove(existingDocument.FilePath) // Best effort deletion, ignore errors
		}
		
		filePath = savePath
	} else {
		// If no new file uploaded, keep the old file path
		filePath = existingDocument.FilePath
	}

	document := entity.Document{
		ID:                 uint(idUint),
		DocumentCategoryID: uint(documentCategoryID),
		Name:               name,
		UserID:             existingDocument.UserID, // Keep original user ID
		FilePath:           filePath,
		Description:        description,
		IsPublished:        isPublished,
	}

	if err := h.usecase.UpdateDocument(id, &document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "document updated successfully", "data": document})
}

func (h *DocumentManagementHandler) DeleteDocument(c *gin.Context) {
	id := c.Param("id")

	// Get document to access file path before deletion
	document, err := h.usecase.GetDocumentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
		return
	}

	// Delete file if it exists
	if document.FilePath != "" {
		_ = os.Remove(document.FilePath) // Best effort deletion, ignore errors
	}

	// Delete record from database
	if err := h.usecase.DeleteDocument(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "document deleted successfully", "data": gin.H{"id": id}})
}
