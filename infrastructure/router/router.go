package router

import (
	"time"

	adminHandler "github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/handler"
	adminRepositoryImpl "github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/repository"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/interface/handler"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/interface/repository"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	// Init repository dan usecase
	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	
	userManagementRepo := adminRepositoryImpl.NewUserManagementRepository(db)
	userManagementUsecase := adminUsecase.NewUserManagementUsecase(userManagementRepo)

	roleManagementRepo := adminRepositoryImpl.NewRoleManagementRepository(db)
	roleManagementUsecase := adminUsecase.NewRoleManagementUsecase(roleManagementRepo)
	
	menuManagementRepo := adminRepositoryImpl.NewMenuManagementRepository(db)
	menuManagementUsecase := adminUsecase.NewMenuManagementUsecase(menuManagementRepo)
	
	devisionManagementRepo := adminRepositoryImpl.NewDivisionManagementRepository(db)
	devisionManagementUsecase := adminUsecase.NewDivisionManagementUsecase(devisionManagementRepo)

	companyManagementRepo := adminRepositoryImpl.NewCompanyManagementRepository(db)
	companyManagementUsecase := adminUsecase.NewCompanyManagementUsecase(companyManagementRepo)

	// Supplier Management
	supplierManagementRepo := adminRepositoryImpl.NewSupplierManagementRepository(db)
	supplierManagementUsecase := adminUsecase.NewSupplierManagementUsecase(supplierManagementRepo)

	// Product Category Management
	productCategoryManagementRepo := adminRepositoryImpl.NewProductCategoryManagementRepository(db)
	productCategoryManagementUsecase := adminUsecase.NewProductCategoryManagementUsecase(productCategoryManagementRepo)

	provinceManagementRepo := adminRepositoryImpl.NewProvinceManagementRepository(db)
	provinceManagementUsecase := adminUsecase.NewProvinceManagementUsecase(provinceManagementRepo)

	cityManagementRepo := adminRepositoryImpl.NewCityManagementRepository(db)
	cityManagementUsecase := adminUsecase.NewCityManagementUsecase(cityManagementRepo)

	// JobPosition Management
	jobPositionManagementRepo := adminRepositoryImpl.NewJobPositionManagementRepository(db)
	jobPositionManagementUsecase := adminUsecase.NewJobPositionManagementUsecase(jobPositionManagementRepo)

	// Business Unit Management
	businessUnitManagementRepo := adminRepositoryImpl.NewBusinessUnitManagementRepository(db)
	businessUnitManagementUsecase := adminUsecase.NewBusinessUnitManagementUsecase(businessUnitManagementRepo)

	// Register handler ke router
	handler.NewAuthHandler(r, authUsecase, db)
	adminHandler.NewUserManagementHandler(r, userManagementUsecase, db)
	adminHandler.NewRoleManagementHandler(r, roleManagementUsecase, db)
	adminHandler.NewMenuManagementHandler(r, menuManagementUsecase, db)
	adminHandler.NewDivisionManagementHandler(r, devisionManagementUsecase, db)
	adminHandler.NewCompanyManagementHandler(r, companyManagementUsecase, db)
	adminHandler.NewProvinceManagementHandler(r, provinceManagementUsecase, db)
	adminHandler.NewCityManagementHandler(r, cityManagementUsecase, db)
	adminHandler.NewJobPositionManagementHandler(r, jobPositionManagementUsecase, db)
	adminHandler.NewBusinessUnitManagementHandler(r, businessUnitManagementUsecase, db)

	// Register Supplier handler
	adminHandler.NewSupplierManagementHandler(r, supplierManagementUsecase, db)

	// Register Product Category handler
	adminHandler.NewProductCategoryManagementHandler(r, productCategoryManagementUsecase)

	// Serve static files for uploaded profile images
	r.Static("/uploads/profile", "./uploads/profile")

	return r
}
