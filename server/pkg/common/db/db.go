package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jquerin/moral-ambiguity/pkg/common/config"
	"github.com/jquerin/moral-ambiguity/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(c *config.Config) *gorm.DB {
	var url string
	// check if running on cloud environment
	cloudUrl := os.Getenv("DATABASE_URL")

	if cloudUrl != "" {
		url = cloudUrl
	} else {
		url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{}, &models.User{})

	return db
}
