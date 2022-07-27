package handler

import (
	"github.com/wagslane/go-rabbitmq"
	"log"
)

type ApiSpecDocHandler struct {
}

func (asdh *ApiSpecDocHandler) Handle(delivery rabbitmq.Delivery) rabbitmq.Action {
	log.Printf("consumed: %v", string(delivery.Body))
	//call process here

	//publish to the required queue success or error
	return rabbitmq.Ack
}

func NewApiSpecDocHandler() Handler {
	return &ApiSpecDocHandler{}
}
