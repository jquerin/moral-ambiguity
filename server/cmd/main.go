package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/jquerin/moral-ambiguity/docs"
	"github.com/jquerin/moral-ambiguity/pkg/common/config"
	"github.com/jquerin/moral-ambiguity/pkg/common/db"
	"github.com/jquerin/moral-ambiguity/pkg/common/routes"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	conf, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	dbHandle := db.InitDB(&conf)

	app := fiber.New()

	routes.RegisterRoutes(app, dbHandle)

	// setup logger middleware
	app.Use(logger.New())
	log.Fatal(app.Listen(config.GetValue("PORT")))
}
