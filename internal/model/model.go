package model

import "time"

type ObjectInfo struct {
	BucketName    string
	Key           string
	LastModified  time.Time
	ContentLength int64
	StorageClass  string
	LastUpdated   time.Time // meta
}

type JobKind string

var JobKindObjectInfo = JobKind("objectInfo")

type Job struct {
	BucketName string
	Key        string
	Interval   int
	JobKind    JobKind
}
