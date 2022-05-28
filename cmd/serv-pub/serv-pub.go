package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/joho/godotenv"
)

var (
	healthy        = true    // This is a simple health flag
	version        = "0.0.1" // App version number, set at build time with -ldflags "-X 'main.version=1.2.3'"
	serviceName    = "users"
	LOCAL_ENV_FILE = "local.env"
	PUBSUB_NAME    = "orderpubsub"
	PUBSUB_TOPIC   = "orders"
)

//
// Main entry point
//
func main() {
	log.Printf("Dapr-starter is starting... service name: %v version: %v", serviceName, version)

	// Load environment information file
	err := godotenv.Load(LOCAL_ENV_FILE)
	if err != nil {
		log.Fatalf("Cannot load local environment information: %s", err)
	}

	serverPort := os.Getenv("DEFAULTPORT")
	log.Printf("serverPort: %s", serverPort)

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	for i := 1; i <= 20; i++ {
		order := `{"orderId":` + strconv.Itoa(i) + `}`

		// Publish an event using Dapr pub/sub
		if err := client.PublishEvent(ctx, PUBSUB_NAME, PUBSUB_TOPIC, []byte(order)); err != nil {
			panic(err)
		}

		log.Printf("Published data: %s", order)

		time.Sleep(2000)
	}

}
