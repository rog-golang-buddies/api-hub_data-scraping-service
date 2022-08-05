package config

//QueueConfig queue configuration
type QueueConfig struct {
	UrlRequestQueue     string `default:"data-scraping-asd"`                  //UrlRequestQueue name to listen to the new events
	ScrapingResultQueue string `default:"storage-update-asd"`                 //Queue name to send processed ApiSpecDoc
	NotificationQueue   string `default:"gateway-scrape_notifications"`       //Queue name to notify a user about error or success (if required)
	Url                 string `default:"amqp://guest:guest@localhost:5672/"` //RabbitMQ url
	Concurrency         int    `default:"30"`                                 //Number of parallel handlers
}
