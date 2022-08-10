package db

import (
	"fmt"
	"log"

	"github.com/jquerin/moral-ambiguity/pkg/common/config"
	"github.com/jquerin/moral-ambiguity/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(c *config.Config) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})

	return db
}
