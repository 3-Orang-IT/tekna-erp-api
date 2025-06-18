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

