package importdata

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, key, contentType *string, data io.Reader) error
}
