package load

import (
	"context"
	"errors"
	"github.com/rog-golang-buddies/internal/model"
)

//ContentLoader loads content by url
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
