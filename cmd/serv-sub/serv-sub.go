package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"

	"github.com/perocha/dapr-starter/config"
)

var sub = &common.Subscription{
	PubsubName: "orderpubsub",
	Topic:      "orders",
	Route:      "/orders",
}

//
// Main entry point
//
func main() {
	// Read configuration file
	cfg, err := config.LoadConfigProvider("serv-sub")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	// Log module name and version
	log.Printf("Dapr-starter module: %s version: %s", cfg.Get("app.name"), cfg.Get("app.version"))

	// Start a new Dapr client
	appPort := cfg.Get("app.port").(string)
	log.Printf("New starting Dapr Subscriber on port %s", appPort)
	s := daprd.NewService(":" + appPort)

	// Register a new subscription
	log.Printf("Subscribing to topic %s", sub.Topic)
	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	fmt.Println("Subscriber received: ", e.Data)
	return false, nil
}
