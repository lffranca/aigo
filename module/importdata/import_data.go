package importdata

import (
	"context"
	"errors"
	"fmt"
	"io"
)

func NewImportData(storage Storage, info FileInfo, generate Generate, basePath *string) (*importData, error) {
	if storage == nil || info == nil || generate == nil || basePath == nil {
		return nil, errors.New("invalid params")
	}

	return &importData{
		storage:  storage,
		info:     info,
		generate: generate,
		basePath: basePath,
	}, nil
}

type importData struct {
	storage  Storage
	info     FileInfo
	generate Generate
	basePath *string
}

func (mod *importData) Import(ctx context.Context, contentType *string, data io.Reader) error {
	key := mod.generate.UUID(ctx)

	path := fmt.Sprintf("%s/%s", *mod.basePath, *key)

	if err := mod.storage.Upload(ctx, &path, contentType, data); err != nil {
		return err
	}

	if err := mod.info.SaveKey(ctx, &path, key); err != nil {
		return err
	}

	return nil
}
