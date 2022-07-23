package load

import (
	"context"
	"errors"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/model"
)

//ContentLoader loads content by url
//go:generate mockgen -source=contentLoader.go -destination=./mocks/contentLoader.go -package=load
type ContentLoader interface {
	Load(ctx context.Context, url string) (*model.FileResource, error)
}

type ContentLoaderImpl struct {
}

func (cl *ContentLoaderImpl) Load(ctx context.Context, url string) (*model.FileResource, error) {
	//load content by url
	return nil, errors.New("not implemented")
}

func NewContentLoader() ContentLoader {
	return &ContentLoaderImpl{}
}
