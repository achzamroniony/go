package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"fullstack-template/backend/internal/model"
)

// APIHandler contains database dependencies for HTTP handlers.
type APIHandler struct {
	db *gorm.DB
}

// NewAPIHandler creates a new instance of APIHandler.
func NewAPIHandler(db *gorm.DB) *APIHandler {
	return &APIHandler{db: db}
}

// HelloHandler handles GET /api/hello by querying topics from PostgreSQL.
func (h *APIHandler) HelloHandler(c *fiber.Ctx) error {
	var topics []model.Topic

	// Fetch all topics from PostgreSQL using GORM
	if err := h.db.Order("id asc").Find(&topics).Error; err != nil {
		response := model.APIResponse{
			Success: false,
			Message: "Gagal memuat data dari database PostgreSQL",
			Errors:  err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Transform database models into string list to match the React frontend model
	var topicList []string
	for _, t := range topics {
		topicList = append(topicList, t.Title)
	}

	// Fallback titles if database is empty
	if len(topicList) == 0 {
		topicList = []string{
			"Golang Basics & Concurrency (Fallback)",
			"Fiber Router & Middleware (Fallback)",
			"PostgreSQL Database Setup (Fallback)",
		}
	}

	data := model.HelloData{
		Message: "Halo dari Backend Go Fiber & PostgreSQL! Koneksi Anda berhasil.",
		Topics:  topicList,
		Version: "1.1.0",
	}

	response := model.APIResponse{
		Success: true,
		Message: "Data berhasil dimuat dari database PostgreSQL!",
		Data:    data,
	}

	return c.Status(http.StatusOK).JSON(response)
}

// HealthHandler handles GET /api/health to check server and database availability.
func (h *APIHandler) HealthHandler(c *fiber.Ctx) error {
	sqlDB, err := h.db.DB()
	dbConnected := false
	if err == nil {
		// Ping database to verify connection is alive
		if pingErr := sqlDB.Ping(); pingErr == nil {
			dbConnected = true
		}
	}

	dbStatus := "connected"
	if !dbConnected {
		dbStatus = "disconnected"
	}

	response := model.APIResponse{
		Success: dbConnected,
		Message: fmt.Sprintf("Server is running. Database status: %s", dbStatus),
	}

	statusCode := http.StatusOK
	if !dbConnected {
		statusCode = http.StatusInternalServerError
	}

	return c.Status(statusCode).JSON(response)
}
