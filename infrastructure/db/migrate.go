package db

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

var models = []interface{}{
	&entity.User{},
	&entity.Role{},
	&entity.Menu{},
	&entity.Company{},
	&entity.City{},
	&entity.Province{},
	&entity.Customer{},
	&entity.Employee{},
	&entity.Division{},
	&entity.JobPosition{},
	&entity.ToDoTemplate{},
	&entity.UnitOfMeasure{},
	&entity.ProductCategory{},
	&entity.BusinessUnit{},
	&entity.Supplier{},
	&entity.Product{},
	&entity.Purchase{},
	&entity.RequestOrder{},
	&entity.SalesOrder{},
	&entity.ProductRequestOrder{},
	&entity.ProductSalesOrder{},
	&entity.TravelCost{},
	&entity.Newsletter{},
	&entity.BankAccount{},
	&entity.ChartOfAccount{},
	&entity.Document{},
	&entity.DocumentCategory{},
}

func DropTables(db *gorm.DB) {
    err := db.Migrator().DropTable(models...)
    if err != nil {
        log.Fatalf("Gagal drop table: %v", err)
    }
	log.Println("Drop tabel selesai")
}

func DropAllTables(db *gorm.DB) {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatalf("Gagal mendapatkan daftar tabel: %v", err)
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			log.Fatalf("Gagal drop tabel %s: %v", table, err)
		}
		log.Printf("Tabel %s berhasil dihapus", table)
	}
	log.Println("Semua tabel berhasil dihapus")
}

func Migrate(db *gorm.DB) error {
    DropAllTables(db) // HATI-HATI: ini akan menghapus data tabel

    err := db.AutoMigrate(models...)
    if err != nil {
        log.Fatalf("Gagal migrate: %v", err)
    }
    log.Println("Migrasi selesai")

    return nil
}

