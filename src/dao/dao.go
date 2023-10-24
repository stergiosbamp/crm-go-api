package dao

import "github.com/stergiosbamp/go-api/src/database"

type DAO interface {
	Get()
}

var db, _ = database.GetDB()
