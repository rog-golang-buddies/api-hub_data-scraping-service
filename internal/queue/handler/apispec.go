package handler

import (
	"encoding/json"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/logger"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/publisher"
	"github.com/wagslane/go-rabbitmq"
)

type ApiSpecDocHandler struct {
	publisher publisher.Publisher
	config    config.QueueConfig
	log       logger.Logger
}

func (asdh *ApiSpecDocHandler) Handle(delivery rabbitmq.Delivery) rabbitmq.Action {
	asdh.log.Infof("consumed: %v", string(delivery.Body))
	//call process here
	var req dto.UrlRequest
	err := json.Unmarshal(delivery.Body, &req)
	if err != nil {
		asdh.log.Errorf("error unmarshalling message: '%v', err: %s", string(delivery.Body), err)
		return rabbitmq.NackDiscard
	}
	//here processing of the request happens...
	asd := apiSpecDoc.ApiSpecDoc{} //TODO replace this stub with process call

	//publish to the required queue success or error
	result := dto.ScrapingResult{IsNotifyUser: req.IsNotifyUser, ApiSpecDoc: asd}
	err = asdh.publish(&delivery, result, asdh.config.ScrapingResultQueue)
	if err != nil {
		asdh.log.Error("error while publishing: ", err)
		//Here is some error while publishing happened - probably something wrong with the queue
		return rabbitmq.NackDiscard
	}
	if req.IsNotifyUser {
		err = asdh.publish(&delivery, dto.NewUserNotification(nil), asdh.config.NotificationQueue)
		if err != nil {
			asdh.log.Error("error while notifying user")
			//don't discard this message because it was published to the storage service successfully
		}
	}
	asdh.log.Info("url scraped successfully")
	return rabbitmq.Ack
}

func (asdh *ApiSpecDocHandler) publish(delivery *rabbitmq.Delivery, message any, queue string) error {
	content, err := json.Marshal(message)
	if err != nil {
		asdh.log.Info("error while marshalling: ", err)
		return err
	}
	return asdh.publisher.Publish(content,
		[]string{queue},
		rabbitmq.WithPublishOptionsCorrelationID(delivery.CorrelationId),
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsPersistentDelivery,
	)
}

func NewApiSpecDocHandler(publisher publisher.Publisher, config config.QueueConfig, log logger.Logger) Handler {
	return &ApiSpecDocHandler{
		publisher: publisher,
		config:    config,
		log:       log,
	}
}
