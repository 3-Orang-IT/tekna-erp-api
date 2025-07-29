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

	// Unit of Measure Management
	unitOfMeasureManagementRepo := adminRepositoryImpl.NewUnitOfMeasureManagementRepository(db)
	unitOfMeasureManagementUsecase := adminUsecase.NewUnitOfMeasureManagementUsecase(unitOfMeasureManagementRepo)

	// Area Management
	areaManagementRepo := adminRepositoryImpl.NewAreaManagementRepository(db)
	areaManagementUsecase := adminUsecase.NewAreaManagementUsecase(areaManagementRepo)

	// Product Management
	productManagementRepo := adminRepositoryImpl.NewProductManagementRepository(db)
	productManagementUsecase := adminUsecase.NewProductManagementUsecase(productManagementRepo)

	// Employee Management
	employeeManagementRepo := adminRepositoryImpl.NewEmployeeManagementRepository(db)
	employeeManagementUsecase := adminUsecase.NewEmployeeManagementUsecase(employeeManagementRepo)

	// Customer Management
	customerManagementRepo := adminRepositoryImpl.NewCustomerManagementRepository(db)
	customerManagementUsecase := adminUsecase.NewCustomerManagementUsecase(customerManagementRepo)

	// Chart of Account Management
	chartOfAccountManagementRepo := adminRepositoryImpl.NewChartOfAccountManagementRepository(db)
	chartOfAccountManagementUsecase := adminUsecase.NewChartOfAccountManagementUsecase(chartOfAccountManagementRepo)
	
	// Bank Account Management
	bankAccountManagementRepo := adminRepositoryImpl.NewBankAccountManagementRepository(db)
	bankAccountManagementUsecase := adminUsecase.NewBankAccountManagementUsecase(bankAccountManagementRepo)
	
	// To Do Template Management
	toDoTemplateManagementRepo := adminRepositoryImpl.NewToDoTemplateManagementRepository(db)
	toDoTemplateManagementUsecase := adminUsecase.NewToDoTemplateManagementUsecase(toDoTemplateManagementRepo)

	// Newsletter Management
	newsletterManagementRepo := adminRepositoryImpl.NewNewsletterManagementRepository(db)
	newsletterManagementUsecase := adminUsecase.NewNewsletterManagementUsecase(newsletterManagementRepo)

	// Register handler ke router
	handler.NewAuthHandler(r, authUsecase, db)
	adminHandler.NewUserManagementHandler(r, userManagementUsecase, db)
	adminHandler.NewRoleManagementHandler(r, roleManagementUsecase, db)
	adminHandler.NewMenuManagementHandler(r, menuManagementUsecase, db)
	adminHandler.NewDivisionManagementHandler(r, devisionManagementUsecase, db)
	adminHandler.NewCompanyManagementHandler(r, companyManagementUsecase, db)
	adminHandler.NewProvinceManagementHandler(r, provinceManagementUsecase, db)
	// Register Newsletter handler
	adminHandler.NewNewsletterManagementHandler(r, newsletterManagementUsecase, db)
	// Register Employee handler
	adminHandler.NewEmployeeManagementHandler(r, employeeManagementUsecase, db)
	adminHandler.NewCityManagementHandler(r, cityManagementUsecase, db)
	adminHandler.NewJobPositionManagementHandler(r, jobPositionManagementUsecase, db)
	adminHandler.NewBusinessUnitManagementHandler(r, businessUnitManagementUsecase, db)
	adminHandler.NewUnitOfMeasureManagementHandler(r, unitOfMeasureManagementUsecase, db)
	adminHandler.NewProductManagementHandler(r, productManagementUsecase, db)
	adminHandler.NewSupplierManagementHandler(r, supplierManagementUsecase, db)
	adminHandler.NewProductCategoryManagementHandler(r, productCategoryManagementUsecase)
	adminHandler.NewCustomerManagementHandler(r, customerManagementUsecase, db)
	adminHandler.NewChartOfAccountManagementHandler(r, chartOfAccountManagementUsecase, db)
	adminHandler.NewBankAccountManagementHandler(r, bankAccountManagementUsecase, db)
	adminHandler.NewToDoTemplateManagementHandler(r, toDoTemplateManagementUsecase, db)
	adminHandler.NewAreaManagementHandler(r, areaManagementUsecase, db)

	// Serve static files for uploads
	r.Static("/uploads/profile", "./uploads/profile")
	r.Static("/uploads/newsletters", "./uploads/newsletters")

	return r
}
