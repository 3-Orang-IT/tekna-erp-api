package config

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"

    "github.com/joho/godotenv"
)

var DB *gorm.DB

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Gagal membaca file .env")
    }
}

func InitDB() *gorm.DB {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("SSL_MODE"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal koneksi ke database:", err)
    }

    DB = db
    return db
}
