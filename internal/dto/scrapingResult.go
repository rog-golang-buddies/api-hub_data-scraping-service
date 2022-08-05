package dto

import "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"

type ScrapingResult struct {
	IsNotifyUser bool

	ApiSpecDoc apiSpecDoc.ApiSpecDoc
}
