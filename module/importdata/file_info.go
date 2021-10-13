package importdata

import "context"

type FileInfo interface {
	SaveKey(ctx context.Context, path, key *string) error
}
