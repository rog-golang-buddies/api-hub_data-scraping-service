package config

//QueueConfig queue configuration
type QueueConfig struct {
	ConsumerQueue string //ConsumerQueue name to listen for the new events
	ProducerQueue string //Queue name to send processed ApiSpecDoc
	ErrorQueue    string //Queue name to notify user about error (if required)
	Url           string //RabbitMQ url
	Concurrency   int    //Number of parallel handlers
}
