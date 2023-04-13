package dao

import "github.com/stergiosbamp/go-api/database"

type DAO interface {
	Get()
}


var db, _ = database.GetDB()
