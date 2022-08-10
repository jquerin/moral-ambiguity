package users

import (
	"github.com/jquerin/moral-ambiguity/pkg/middleware"

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

	userRoutes := app.Group("/api/users")
	userRoutes.Get("/:id", dbHandle.GetUserById)
	userRoutes.Post("/", dbHandle.AddUser)
	userRoutes.Patch("/:id", middleware.Protected(), dbHandle.UpdateUser)
	userRoutes.Delete("/:id", middleware.Protected(), dbHandle.DeleteUser)
}
