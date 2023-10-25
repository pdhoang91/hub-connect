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
	PGURL         = GetString("PGURL", "host=telecomdb port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	MaxOpenConns  = GetInt("MaxOpenConns", 25)
	MaxIdleConns  = GetInt("MaxIdleConns", 20)
	HTTPPort      = GetString("HTTPPort", ":80")
	SwaggerDomain = GetString("SwaggerDomain", "127.0.0.1")
)

// Config represents the application configuration.
type Config struct {
	HTTPPort      string
	PGURL         string
	SwaggerDomain string
	MaxOpenConns  int
	MaxIdleConns  int
}

// NewConfig returns an initialized Config based on environment variables.
func NewConfig() *Config {

	// Initialize the Config struct with environment variable values
	cfg := Config{
		HTTPPort:      GetString("HTTPPort", ":80"),
		MaxOpenConns:  GetInt("MaxOpenConns", 25), // Default value, suitable for Core i7 CPU, 8GB RAM.
		MaxIdleConns:  GetInt("MaxIdleConns", 20), //Default value, suitable for Core i7 CPU, 8GB RAM.
		PGURL:         GetString("PGURL", "host=telecomdb port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"),
		SwaggerDomain: GetString("SwaggerDomain", "127.0.0.1"),
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
			sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
			sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

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
