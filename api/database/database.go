package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/stergiosbamp/go-api/config"
)

func GetDB() (*gorm.DB, error) {

	var config config.Config
	dsn := config.CreateDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// Maybe AutoMigrate here...

	return db, err
}
