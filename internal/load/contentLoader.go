package load

import (
	"context"
	"errors"

	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
)

//ContentLoader loads content by url
//go:generate mockgen -source=contentLoader.go -destination=./mocks/contentLoader.go -package=load
type ContentLoader interface {
	Load(ctx context.Context, url string) (*fileresource.FileResource, error)
}

type ContentLoaderImpl struct {
}

// Gets context and an url of a OpenApi file (Swagger file) string as parameter and returns a FileResource containing the link, optionally name and main content of the file.
func (cl *ContentLoaderImpl) Load(ctx context.Context, url string) (*fileresource.FileResource, error) {
	//load content by url
	return nil, errors.New("not implemented")
}

func NewContentLoader() ContentLoader {
	return &ContentLoaderImpl{}
}
