package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DbHandler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	dbHandle := &DbHandler{
		DB: db,
	}

	routes := app.Group("/api/auth")
	routes.Post("/login", dbHandle.Login)
}
