package service

import (
	"context"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/s3-metrics-exporter/internal/model"
)

type s3client interface {
	GetFileInfo(ctx context.Context, bucketName string, key string) (objInfo model.ObjectInfo, err error)
}

type Fetcher struct {
	S3Client s3client
}

func (f *Fetcher) Fetch() (err error) {
	jobs := map[string]model.Job{}
	_, err = toml.DecodeFile("./deployments/config.toml", &jobs)
	if err != nil {
		return err
	}

	for i, v := range jobs {
		fmt.Println(i, v)
	}
	return nil
}
