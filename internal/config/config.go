package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DEPLOYMENT int

const (
	PRODUCTION DEPLOYMENT = iota
	DEVELOPMENT
)

type Config struct {
	Port   int
	DB_URI string
	Env    DEPLOYMENT
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("warning: .env file not found: %v", err)
	}

	port := 8080 // default
	if portStr := os.Getenv("CODEBIN_PORT"); portStr != "" {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			log.Printf("warning: could not parse CODEBIN_PORT=%q: %v (defaulting to 8080)", portStr, err)
		} else {
			port = p
		}
	}

	dbURI := os.Getenv("CODEBIN_DATA_PATH")
	if dbURI == "" {
		return nil, fmt.Errorf("no URI provided for database (CODEBIN_DATA_PATH not set)")
	}

	deployment := os.Getenv("ENVIRONMENT")
	depType := PRODUCTION

	switch deployment {
	case "":
		log.Printf("warning: ENVIRONMENT not set, defaulting to production")
	case "development":
		depType = DEVELOPMENT
	case "production":
		depType = PRODUCTION
	default:
		log.Printf("warning: unknown ENVIRONMENT value %q, defaulting to production", deployment)
	}

	cfg := &Config{
		Port:   port,
		DB_URI: dbURI,
		Env:    depType,
	}

	return cfg, nil
}
