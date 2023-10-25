// Package config provides configuration settings for the application.
package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
// PGURL         = GetString("PGURL", "host=telecomdb port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
// MaxOpenConns  = GetInt("MaxOpenConns", 25)
// MaxIdleConns  = GetInt("MaxIdleConns", 20)
// HTTPPort      = GetString("HTTPPort", ":80")
// SwaggerDomain = GetString("SwaggerDomain", "127.0.0.1")
)

// Config represents the application configuration.
type Config struct {
	APP_PORT       string
	PGURL          string
	SWAGGER_DOMAIN string
	MAX_OPEN_CONNS int
	MAX_IDLE_CONNS int
}

// NewConfig returns an initialized Config based on environment variables.
func NewConfig() *Config {

	// Initialize the Config struct with environment variable values
	cfg := Config{
		APP_PORT:       GetString("APP_PORT", "80"),
		MAX_OPEN_CONNS: GetInt("MAX_OPEN_CONNS", 25), // Default value, suitable for Core i7 CPU, 8GB RAM.
		MAX_IDLE_CONNS: GetInt("MAX_IDLE_CONNS", 20), //Default value, suitable for Core i7 CPU, 8GB RAM.
		PGURL:          GetString("PGURL", "host=telecomdb port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"),
		SWAGGER_DOMAIN: GetString("SWAGGER_DOMAIN", "127.0.0.1"),
	}
	return &cfg
}

// InitDBConnection initializes a database connection using the provided configuration.
func InitDBConnection(cfg *Config) (*gorm.DB, error) {
	maxRetries := 5

	for retry := 1; retry <= maxRetries; retry++ {
		db, err := gorm.Open(postgres.Open(cfg.PGURL), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			sqlDB, err := db.DB()
			if err != nil {
				return nil, err
			}

			// Set maximum idle and in-use connections
			sqlDB.SetMaxIdleConns(cfg.MAX_IDLE_CONNS)
			sqlDB.SetMaxOpenConns(cfg.MAX_OPEN_CONNS)

			fmt.Println("Connected to the database!")
			return db, nil
		}

		fmt.Printf("Attempt %d failed: %v\n", retry, err)
		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to the database after %d retries", maxRetries)
}

func CloseDBConnection(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		return err
	}
	return database.Close()
}
