package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/perocha/serv-sub/config"
)

var sub = &common.Subscription{
	PubsubName: "orderpubsub",
	Topic:      "orders",
	Route:      "/orders",
}

func Run(cfg *config.Config) {
	// Log module name and version
	log.Printf("Dapr-starter module: %s version: %s", cfg.Name, cfg.Version)

	// Start a new Dapr client
	log.Printf("New starting Dapr Subscriber on port %s", cfg.App.Port)
	s := daprd.NewService(":" + cfg.App.Port)

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
