package main

import (
	"context"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler"
	"log"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conf := config.ReadConfig()

	handl := handler.NewApiSpecDocHandler()
	listener := queue.NewListener(&conf.QueueConfig, handl)

	err := listener.Start(ctx)

	if err != nil {
		log.Println("error while listening queue ", err)
		cancel()
		os.Exit(1)
		return
	}
	log.Println("application stopped gracefully (not)")
}
