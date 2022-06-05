package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"

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
	log.Printf("APP_ID: %s", os.Getenv("APP_ID"))
	log.Printf("APP_PORT: %s", os.Getenv("APP_PORT"))
	log.Printf("DAPR_HTTP_PORT: %s", os.Getenv("DAPR_HTTP_PORT"))
	log.Printf("DAPR_GRPC_PORT: %s", os.Getenv("DAPR_GRPC_PORT"))
	log.Printf("NAMESPACE: %s", os.Getenv("NAMESPACE"))

	// Use viper to read config from environment variables
	viper.SetEnvPrefix("dapr")
	viper.AutomaticEnv()
	viper.SetDefault("port", "8080")
	viper.SetDefault("pubsub_name", "orderpubsub")
	viper.SetDefault("topic", "orders")
	viper.SetDefault("route", "/orders")
	viper.SetDefault("app_id", "order-sub")
	viper.SetDefault("app_port", "6001")

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
