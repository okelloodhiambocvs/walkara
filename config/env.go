package config

import "os"

type Config struct {
	Port string
	DB   string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := os.Getenv("DB_PATH")
	if db == "" {
		db = "./walkara.db"
	}

	return &Config{
		Port: port,
		DB:   db,
	}
}