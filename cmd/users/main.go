package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"
)

var (
	healthy     = true    // This is a simple health flag
	version     = "0.0.1" // App version number, set at build time with -ldflags "-X 'main.version=1.2.3'"
	serviceName = "users"
	defaultPort = 8080
)

//
// Main entry point
//
func main() {
	log.Printf("dapr-starter is starting... service name: %v version: %v", serviceName, version)

	// Port to listen on
	serverPort := os.Getenv("PORT")
	log.Printf("serverPort: %s", serverPort)

	// Use gorilla/mux for routing
	router := mux.NewRouter()
	if router != nil {
		panic("Error initializing")
	}

}
