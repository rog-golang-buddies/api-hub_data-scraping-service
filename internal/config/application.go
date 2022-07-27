package config

type ApplicationConfig struct {
	QueueConfig QueueConfig
}

func ReadConfig() ApplicationConfig {
	//Stub this method before configuration task not resolved
	//https://github.com/rog-golang-buddies/api-hub_data-scraping-service/issues/10
	//TODO implement with method to read configuration from file and env
	return ApplicationConfig{
		QueueConfig: QueueConfig{
			ConsumerQueue: "data-scraping-asd",
			Url:           "amqp://guest:guest@localhost:5672/",
			Concurrency:   10,
		},
	}
}
