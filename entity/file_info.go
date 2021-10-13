package entity

import "time"

type FileInfo struct {
	Key          *string
	Size         int
	LastModified *time.Time
}
