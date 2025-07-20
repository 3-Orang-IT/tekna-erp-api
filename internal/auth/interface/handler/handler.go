package handler

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/middleware"
	usecase "github.com/3-Orang-IT/tekna-erp-api/internal/auth/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/utils"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
    usecase usecase.AuthUsecase
}

type MenuNode struct {
    ID       uint       `json:"id"`
    ParentID *uint      `json:"parent_id"`
    ModulID  uint       `json:"modul_id"`
    Name     string     `json:"name"`
    URL      string     `json:"path,omitempty"`
    Icon     string     `json:"icon"`
    Order    int        `json:"order"`
    Children []MenuNode `json:"children,omitempty"`
}

func NewAuthHandler(r *gin.Engine, uc usecase.AuthUsecase, db *gorm.DB) {
    h := &AuthHandler{uc}
    api := r.Group("/api/v1")
    api.POST("/auth/register", h.Register)
    api.POST("/auth/login", h.Login)

    protected := api.Group("/")
    protected.Use(middleware.JWTAuthMiddleware(db))
    protected.GET("/menus", h.GetMenus)
}

func (h *AuthHandler) Register(c *gin.Context) {
    var user entity.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.usecase.Register(&user); err != nil {
        if err.Error() == "email already registered" {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var input struct {
        Username    string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.usecase.Login(input.Username, input.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
        return
    }

    roles := []string{}
    for _, role := range user.Role {
        roles = append(roles, role.Name)
    }

    c.JSON(http.StatusOK, gin.H{
        "data" : gin.H{
            "token": token,
            "user":  gin.H{
                "id":    user.ID,
                "name":  user.Name,
                "email": user.Email,
                "role": roles,
            },
        },
    })
}

func (h *AuthHandler) GetMenus(c *gin.Context) {
    // Ambil user ID dari context token
    userIdInterface, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
        return
    }

    userID, ok := userIdInterface.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
        return
    }

    // Ambil menu dari usecase
    menus, err := h.usecase.GetMenus(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Build tree function
    var buildMenuTree func([]entity.Menu, *uint) []MenuNode
    buildMenuTree = func(menus []entity.Menu, parentID *uint) []MenuNode {
        var nodes []MenuNode
        var filteredMenus []entity.Menu

        // Filter menu dengan ParentID sesuai
        for _, menu := range menus {
            if (menu.ParentID == nil && parentID == nil) ||
                (menu.ParentID != nil && parentID != nil && *menu.ParentID == *parentID) {
                filteredMenus = append(filteredMenus, menu)
            }
        }

        // Urutkan berdasarkan Order
        sort.SliceStable(filteredMenus, func(i, j int) bool {
            return filteredMenus[i].Order < filteredMenus[j].Order
        })

        for _, menu := range filteredMenus {
            children := buildMenuTree(menus, &menu.ID)

            node := MenuNode{
                ID:       menu.ID,
                ParentID: menu.ParentID,
                ModulID:  menu.ModulID,
                Name:     menu.Name,
                Icon:     menu.Icon,
                Order:    menu.Order,
                URL:      menu.URL,
            }

            // Jika punya children, kosongkan URL dan tambahkan children
            if len(children) > 0 {
                node.URL = ""
                node.Children = children
            }

            nodes = append(nodes, node)
        }

        return nodes
    }

    menuTree := buildMenuTree(menus, nil)
    c.JSON(http.StatusOK, gin.H{"data": menuTree})
}



func parseUint(str string) uint {
    var i uint
    fmt.Sscanf(str, "%d", &i)
    return i
}
