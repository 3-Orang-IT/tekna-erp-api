package main

import (
	"fmt"
	"log"
	"os"

	"github.com/3-Orang-IT/tekna-erp-api/infrastructure/db"
	"github.com/3-Orang-IT/tekna-erp-api/infrastructure/router"
	"github.com/3-Orang-IT/tekna-erp-api/infrastructure/seeders"
	"github.com/3-Orang-IT/tekna-erp-api/internal/config"
)

func main() {
    config.LoadEnv()
    database := config.InitDB()

    if err := db.Migrate(database); err != nil {
        log.Fatal("Migrasi gagal:", err)
    }

    seeders.Seed(database)

    r := router.InitRoutes(database)

    host := os.Getenv("HOST")
    port := os.Getenv("PORT")
    if host == "" {
        host = "localhost" // Default value
    }
    if port == "" {
        port = "8080" // Default value
    }

    address := fmt.Sprintf("%s:%s", host, port)
    log.Printf("Server siap berjalan di http://%s", address)
    if err := r.Run(address); err != nil {
        log.Fatal("Gagal menjalankan server:", err)
    }
}
