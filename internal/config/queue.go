package config

// QueueConfig queue configuration
type QueueConfig struct {
	//UrlRequestQueue name to listen to the new events
	UrlRequestQueue string `default:"data-scraping-asd" envconfig:"URL_REQUEST_QUEUE"`
	//ScrapingResultQueue represents a queue name to send processed ApiSpecDoc
	ScrapingResultQueue string `default:"storage-update-asd" envconfig:"SCRAPING_RESULT_QUEUE"`
	//NotificationQueue represents a queue name to notify a user about an error or success (if required)
	NotificationQueue string `default:"gateway-scrape-notifications" envconfig:"NOTIFICATION_QUEUE"`
	//Url is a RabbitMQ url
	Url string `default:"amqp://guest:guest@localhost:5672/"`
	//Concurrency represents number of parallel handlers
	Concurrency int `default:"30"`
}
