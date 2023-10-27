package dao

import (
	"github.com/stergiosbamp/go-api/src/database"
	"gorm.io/gorm"
)


var db *gorm.DB

func init() {
	db = database.DB
}

