package database

import (
	"log"

	"github.com/stergiosbamp/go-api/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {

	var config config.Config
	dsn := config.CreateDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db, err
}
