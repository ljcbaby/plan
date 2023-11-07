package main

import (
	"fmt"
	"log"

	"github.com/ljcbaby/plan/config"
	"github.com/ljcbaby/plan/database"
	"github.com/ljcbaby/plan/router"
)

func main() {
	// Load config
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Init database
	err = database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect MySQL: %v", err)
	}

	// Setup router
	r := router.SetupRouter()

	// Start server
	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	err = r.Run(addr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
