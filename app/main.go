// @title Integration Auth API
// @version 1.0
// @description API สำหรับ OAuth Token
// @host localhost:8080
package main

import (
	"integration-auth-service/configs"
	"integration-auth-service/modules/servers"
	databases "integration-auth-service/pkg/databases"
	"integration-auth-service/pkg/loggers"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
)

func main() {
	println("Starting Integration Auth Service...")
	config := loadConfig()
	db := initDatabase(config)
	defer db.Close()
	cache := initCache()
	logger := initLogger(db)

	println("Starting server...")
	server := servers.NewServer(&config, db, cache, logger)
	server.Start()
}

func loadConfig() configs.Configs {
	// Load environment variables for Localhost development
	// if err := godotenv.Load("../.env"); err != nil {

	// Load dotenv config
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}

	config := configs.LoadEnv()

	return config
}

func initDatabase(config configs.Configs) *sqlx.DB {
	// New Database
	db, err := databases.NewPostgreSqlDbConnection(&config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}

func initCache() *cache.Cache {
	// Initialize cache with a default expiration time and cleanup interval
	c := cache.New(5*time.Minute, 10*time.Minute)
	return c
}

func initLogger(db *sqlx.DB) *loggers.Logger {
	// Initialize logger
	logger := loggers.NewLogger(db)
	return &logger
}
