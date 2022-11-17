package model

import "encoding/json"

type FileH struct {
	Data json.RawMessage
	Key  string `json:"key"`
}

type File struct {
	BucketName string `json:"bucket"`
	FileName   string `json:"file"`
}
