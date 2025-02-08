package main

import (
	"fmt"
	"log"

	"github.com/abyanmajid/matcha/env"
)

type EnvConfig struct {
	Port     int    `name:"PORT" required:"true" default:"8080"`
	Debug    bool   `name:"DEBUG" required:"false" default:"false"`
	Database string `name:"DATABASE_URL" required:"true"`
}

func main() {
	if err := env.Dotenv(".env"); err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	}

	var config EnvConfig
	if err := env.Load(&config); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Config: %+v\n", config)
}
