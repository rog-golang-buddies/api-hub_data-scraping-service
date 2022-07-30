package publisher

import (
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/wagslane/go-rabbitmq"
	"io"
	"log"
)

//Publisher is just an interface for the library publisher which doesn't have one.
type Publisher interface {
	io.Closer
	Publish(
		data []byte,
		routingKeys []string,
		optionFuncs ...func(*rabbitmq.PublishOptions),
	) error
}

func NewPublisher(conf config.QueueConfig) (Publisher, error) {
	return rabbitmq.NewPublisher(
		conf.Url,
		rabbitmq.Config{},
	)
}

func ClosePublisher(publisher Publisher) {
	log.Println("closing publisher")
	err := publisher.Close()
	if err != nil {
		log.Println("error while closing publisher: ", err)
	}
}
