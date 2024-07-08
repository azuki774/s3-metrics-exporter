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
	S3Client   s3client
	ConfigFile string
}

func (f *Fetcher) Fetch() (err error) {
	jobs := map[string]model.Job{}
	_, err = toml.DecodeFile(f.ConfigFile, &jobs)
	if err != nil {
		return err
	}

	for _, v := range jobs {
		job := model.Job{
			BucketName: v.BucketName,
			Key:        v.Key,
			Interval:   v.Interval,
			JobKind:    v.JobKind,
		}

		objInfo, err := f.S3Client.GetFileInfo(context.TODO(), job.BucketName, job.Key)
		if err != nil {
			return err
		}

		fmt.Printf("(%%#v) %#v\n", objInfo)
	}
	return nil
}
