package config

//QueueConfig queue configuration
type QueueConfig struct {
	UrlRequestQueue     string //UrlRequestQueue name to listen for the new events
	ScrapingResultQueue string //Queue name to send processed ApiSpecDoc
	NotificationQueue   string //Queue name to notify user about error or success (if required)
	Url                 string //RabbitMQ url
	Concurrency         int    //Number of parallel handlers
}
