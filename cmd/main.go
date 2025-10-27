package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gitnishantjha/Api_Gateway/pkg/config"
	"github.com/gitnishantjha/Api_Gateway/pkg/router"
)

func main() {

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	mainRouter, err := router.NewRouter(cfg)
	if err != nil {
		log.Fatalf("Error creating router: %v", err)
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting Porta API Gateway on %s", addr)

	if err := http.ListenAndServe(addr, mainRouter); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

}
