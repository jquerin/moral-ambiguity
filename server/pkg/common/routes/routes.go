package routes

import (
	"github.com/jquerin/moral-ambiguity/pkg/auth"
	"github.com/jquerin/moral-ambiguity/pkg/products"
	"github.com/jquerin/moral-ambiguity/pkg/users"

	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	// base API url
	app.Group("/api", logger.New())

	// products API
	products.RegisterRoutes(app, db)

	// auth API
	auth.RegisterRoutes(app, db)

	// users API
	users.RegisterRoutes(app, db)

	// add Swagger docs route
	app.Get("/swagger/*", swagger.HandlerDefault)
}
