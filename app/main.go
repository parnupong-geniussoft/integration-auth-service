package main

import (
	"integration-auth-service/configs"
	"integration-auth-service/modules/servers"
	databases "integration-auth-service/pkg/databases"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
)

func main() {
	// Load dotenv config
	if err := godotenv.Load("../.env"); err != nil {
		panic(err.Error())
	}

	cfg := configs.LoadEnv()

	// New Database
	db, err := databases.NewPostgreSQLDBConnection(&cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()

	c := cache.New(5*time.Minute, 10*time.Minute)

	s := servers.NewServer(&cfg, db, c)
	s.Start()
}
