package importdata

import "context"

type Generate interface {
	UUID(ctx context.Context) *string
}
