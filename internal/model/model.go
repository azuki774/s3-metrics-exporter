package model

import "time"

type ObjectInfo struct {
	BucketName    string
	Key           string
	LastModified  time.Time
	ContentLength int64
	LastUpdated   time.Time // meta
}
