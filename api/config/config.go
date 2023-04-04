package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbName string
}

func (config *Config) CreateDSN() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env")
	}

	user := os.Getenv("DB_USER")
	config.DbUser = user

	pass := os.Getenv("DB_PASS")
	config.DbPass = pass

	host := os.Getenv("DB_HOST")
	config.DbHost = host

	port := os.Getenv("DB_PORT")
	config.DbPort = port

	dbName := os.Getenv("DB_NAME")
	config.DbName = dbName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", 
						config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName)

	return dsn
}
