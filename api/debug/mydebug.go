package main

import (
	"fmt"
	// "time"

	"github.com/stergiosbamp/go-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:pass@tcp(0.0.0.0:3306)/go-api?charset=utf8&parseTime=true&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	res := db.Model(models.Customer{})	

	fmt.Println(res)
}
