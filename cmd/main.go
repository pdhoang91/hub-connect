package main

import (
	"fmt"
	"hub-connect/config"
	"hub-connect/internal/app"
)

func main() {

	cfg := config.NewConfig()

	db, err := config.InitDBConnection(cfg)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer config.CloseDBConnection(db)

	app := app.InitializeHTTPServer(cfg, db)
	app.Run(cfg.HTTPPort)
}
