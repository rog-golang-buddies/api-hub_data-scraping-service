package main

import (
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler"
	"log"
)

func main() {
	conf := config.ReadConfig()

	handl := handler.NewApiSpecDocHandler()
	listener := queue.NewListener(&conf.QueueConfig, handl)

	err := listener.Start()
	if err != nil {
		log.Println("error while listening queue ", err)
		return
	}
	log.Println("application stopped gracefully (not)")
}
