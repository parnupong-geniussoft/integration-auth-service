package configs

import "os"

type Configs struct {
	App        Fiber
	PostgreSQL PostgreSQL
	Auth       Auth
}

type Fiber struct {
	Host string
	Port string
}

// Database
type PostgreSQL struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Database string
	SSLMode  string
}

// Auth
type Auth struct {
	OauthJwtSecret string
}

func LoadEnv() Configs {
	cfg := Configs{}

	// Fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	// Auth Configs
	cfg.Auth.OauthJwtSecret = os.Getenv("OAUTH_JWT_SECRET")

	return cfg
}
