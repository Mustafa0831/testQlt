package server

import "qlt/internal/database"

// Config ...
type Config struct {
	Database *database.Config
}

// NewConfig generates configurations for the Server
func NewConfig() *Config {
	return &Config{
		Database: database.NewConfig(),
	}
}
