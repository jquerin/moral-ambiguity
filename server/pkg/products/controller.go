package products

import (
	"github.com/jquerin/moral-ambiguity/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	dbHandle := &handler{
		DB: db,
	}

	routes := app.Group("/api/products")
	routes.Get("/", dbHandle.GetAllProducts)
	routes.Get("/:id", dbHandle.GetProductByID)
	routes.Post("/", dbHandle.AddProduct)
	routes.Put("/:id", middleware.Protected(), dbHandle.UpdateProduct)
	routes.Delete("/:id", middleware.Protected(), dbHandle.DeleteProduct)
}
