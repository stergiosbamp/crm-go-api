package database

import (
	"log"

	_ "github.com/joho/godotenv/autoload" // Load .env file automatically
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/stergiosbamp/go-api/src/config"
	"github.com/stergiosbamp/go-api/src/models"
)

var DB *gorm.DB

func init() {
	configuration := config.NewConfig()
	
	dsn := configuration.CreateDSN()
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db
}

func Migrate() error {
	err := DB.AutoMigrate(&models.Address{}, &models.Contact{}, &models.Customer{}, &models.User{})
	if err != nil {
		return err
	}

	return nil
}
