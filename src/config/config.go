package config

import (
	"fmt"
	"os"
)

type Config struct {
	DbUser         string
	DbPass         string
	DbHost         string
	DbPort         string
	DbName         string
	RedisHost      string
	RedisPort      string
	RedisPass      string
	RedisDB        string
	Authentication string
	SecretKey      string
}

func NewConfig() *Config {
	var config = new(Config)

	config.DbUser = os.Getenv("DB_USER")
	config.DbPass = os.Getenv("DB_PASS")
	config.DbHost = os.Getenv("DB_HOST")
	config.DbPort = os.Getenv("DB_PORT")
	config.DbName = os.Getenv("DB_NAME")

	config.RedisHost = os.Getenv("REDIS_HOST")
	config.RedisPort = os.Getenv("REDIS_PORT")
	config.RedisPass = os.Getenv("REDIS_PASS")
	config.RedisDB = os.Getenv("REDIS_DB")

	config.SecretKey = os.Getenv("SECRET_KEY")
	config.Authentication = os.Getenv("AUTHENTICATION")

	return config
}

func (config *Config) CreateDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName)

	return dsn
}

func (config *Config) GetSecretKey() string {
	return config.SecretKey
}

func (config *Config) UseAuth() bool {
	return config.Authentication == "true"
}
