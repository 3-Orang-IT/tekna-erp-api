package db

import (
	"log"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/domain/entity"
	"gorm.io/gorm"
)

var models = []interface{}{
	&entity.User{},
	&entity.Role{},
	&entity.Menu{},
}

func DropTables(db *gorm.DB) {
    err := db.Migrator().DropTable(models...)
    if err != nil {
        log.Fatalf("Gagal drop table: %v", err)
    }
	log.Println("Drop tabel selesai")
}

func Migrate(db *gorm.DB) error {
    DropTables(db) // HATI-HATI: ini akan menghapus data tabel

    err := db.AutoMigrate(models...)
    if err != nil {
        log.Fatalf("Gagal migrate: %v", err)
    }
    log.Println("Migrasi selesai")

    return nil
}

