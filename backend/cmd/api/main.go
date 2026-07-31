package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"

	"fullstack-template/backend/config"
	"fullstack-template/backend/internal/database"
	"fullstack-template/backend/internal/handler"
	"fullstack-template/backend/internal/middleware"
	"fullstack-template/backend/internal/model"
)

func main() {
	// 1. Load Application Configuration
	cfg := config.LoadConfig()

	// 2. Initialize Database Connection
	db := database.InitDB(cfg)

	// 3. Run GORM Auto Migrations
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(&model.Topic{}); err != nil {
		log.Fatalf("Error running database migrations: %v", err)
	}

	// 4. Seed Default Topics if database is empty
	seedTopics(db)

	// 5. Initialize Handlers with Dependency Injection
	h := handler.NewAPIHandler(db)

	// 6. Initialize Fiber App
	app := fiber.New(fiber.Config{
		AppName: "Go Fiber React Learning Portal API v1.1",
	})

	// 7. Register Standard Middlewares
	app.Use(recover.New()) // Recovers from panics anywhere in the stack
	app.Use(logger.New())  // Logs HTTP request details

	// 8. Setup Custom CORS Middleware
	middleware.SetupCORS(app)

	// 9. Register API Routing
	api := app.Group("/api")
	api.Get("/health", h.HealthHandler)
	api.Get("/hello", h.HelloHandler)

	// 10. Start the Server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting %s server in %s mode on %s", app.Config().AppName, cfg.Env, addr)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// seedTopics inserts initial default syllabus topics if none exist.
func seedTopics(db *gorm.DB) {
	var count int64
	db.Model(&model.Topic{}).Count(&count)
	if count == 0 {
		log.Println("Seeding default learning topics to PostgreSQL database...")
		defaultTopics := []model.Topic{
			{Title: "Golang Basics & Concurrency", Description: "Sintaksis dasar Go, pointer, struct, interface, goroutine, dan channel."},
			{Title: "Fiber Router & Middleware", Description: "Membangun REST API berkinerja tinggi menggunakan routing dan middleware Fiber."},
			{Title: "REST API Design (Clean Architecture)", Description: "Menyusun struktur kode Go secara profesional berbasis clean architecture."},
			{Title: "PostgreSQL & GORM Integration", Description: "Koneksi database relasional, pemodelan skema, migrasi, dan query ORM dengan GORM."},
			{Title: "React Hooks & Context (Vite + TS)", Description: "Membangun frontend modular dengan state management global dan pengetikan tipe aman."},
			{Title: "API Integration & CORS Policy", Description: "Menghubungkan klien React dengan backend API Go secara aman lintas origin."},
		}

		if err := db.Create(&defaultTopics).Error; err != nil {
			log.Printf("Warning: Failed to seed default topics: %v", err)
		} else {
			log.Println("Database seeded successfully with default syllabus topics!")
		}
	}
}
