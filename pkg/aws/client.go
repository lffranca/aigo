package aws

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lffranca/aigo/entity"
	"io"
)

func New(bucket *string) (*client, error) {
	if bucket == nil {
		return nil, errors.New("invalid params")
	}

	return &client{
		bucket: bucket,
	}, nil
}

type client struct {
	bucket *string
}

func (pkg *client) PreSign(ctx context.Context, key, contentType *string) (*string, error) {
	client, err := pkg.clientS3(ctx)
	if err != nil {
		return nil, err
	}

	psClient := s3.NewPresignClient(client)
	res, err := psClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket:              pkg.bucket,
		Key:                 key,
		ResponseContentType: contentType,
	})
	if err != nil {
		return nil, err
	}

	return &res.URL, nil
}

func (pkg *client) Upload(ctx context.Context, key, contentType *string, data io.Reader) error {
	client, err := pkg.clientS3(ctx)
	if err != nil {
		return err
	}

	uploader := manager.NewUploader(client)
	if _, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Key:         key,
		ContentType: contentType,
		Body:        data,
		Bucket:      pkg.bucket,
	}); err != nil {
		return err
	}

	return nil
}

func (pkg *client) ListObjects(ctx context.Context) ([]*entity.FileInfo, error) {
	client, err := pkg.clientS3(ctx)
	if err != nil {
		return nil, err
	}

	output, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: pkg.bucket,
	})
	if err != nil {
		return nil, err
	}

	var files []*entity.FileInfo
	for _, object := range output.Contents {
		files = append(files, &entity.FileInfo{
			Key:          object.Key,
			Size:         int(object.Size),
			LastModified: object.LastModified,
		})
	}

	return files, nil
}

func (pkg *client) clientS3(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}
