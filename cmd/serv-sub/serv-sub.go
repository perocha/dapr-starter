package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

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
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	appPort := cfg.AppPort
	log.Printf("cfg.AppPort %s", appPort)

	log.Printf("APP_ID: %s", os.Getenv("APP_ID"))
	log.Printf("APP_PORT: %s", os.Getenv("APP_PORT"))
	log.Printf("DAPR_HTTP_PORT: %s", os.Getenv("DAPR_HTTP_PORT"))
	log.Printf("DAPR_GRPC_PORT: %s", os.Getenv("DAPR_GRPC_PORT"))
	log.Printf("NAMESPACE: %s", os.Getenv("NAMESPACE"))

	/*	if !isSet {
			log.Fatalf("APP_PORT is not set")
		}
	*/
	//	appPort := "6001"
	log.Printf("New starting Dapr Subscriber on port %s", appPort)

	s := daprd.NewService(":" + appPort)
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
