package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"fullstack-template/backend/config"
)

// InitDB initializes GORM PostgreSQL database connection.
func InitDB(cfg *config.Config) *gorm.DB {
	// Construct the PostgreSQL DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	log.Printf("Connecting to PostgreSQL database '%s' on %s:%s...", cfg.DBName, cfg.DBHost, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Logs SQL statements
	})

	if err != nil {
		log.Println("\n==============================================================")
		log.Printf("⚠️ ERROR: Gagal terhubung ke database PostgreSQL: %v\n", err)
		log.Println("PANDUAN PENYELASAIAN:")
		log.Printf("1. Pastikan server PostgreSQL Anda aktif di host '%s' dan port '%s'.\n", cfg.DBHost, cfg.DBPort)
		log.Printf("2. Pastikan database bernama '%s' sudah dibuat.\n", cfg.DBName)
		log.Printf("3. Periksa kembali user '%s' dan password di file '.env'.\n", cfg.DBUser)
		log.Println("==============================================================")
		log.Fatalf("Database connection critical failure.")
	}

	log.Println("✅ Database PostgreSQL berhasil terhubung!")
	return db
}
