package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
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
	// Read APP_PORT from container
	//	appPort, isSet := os.LookupEnv("APP_PORT")
	appPort := os.Getenv("APP_PORT")
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
