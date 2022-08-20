package dto

// UrlRequest represents listening request model
type UrlRequest struct {
	//File url to scrape data
	FileUrl string `json:"file_url"`

	//A flag is a notification required related to an error notification in case of an error
	//Notification is required when this is the request from the user and doesn't require it
	//if it is the request from the storage and update service.
	IsNotifyUser bool `json:"is_notify_user"`
}
