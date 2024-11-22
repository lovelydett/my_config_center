package db

import (
	"context"
	"io"
	"time"
	"wolf/drivers"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var pg = drivers.Pg
var redis = drivers.Redis
var bucket = drivers.Bucket

func InsertChunkMeta(filename string, tags []string, description string) error {
	params := map[string]interface{}{
		"filename":    filename,
		"tags":        tags,
		"description": description,
	}
	_, err := pg.Exec("INSERT INTO files (filename, tags, description) VALUES (:filename, :tags, :description)", params)
	if err != nil {
		return err
	}
	return nil
}

func InitChunkUpload(key string) bool {
	val, _ := redis.SetNX(context.Background(), key, "inited", time.Minute).Result()
	if val {
		return true
	}
	return false
}

func UpdateChunkUploadImur(key string, objectId string) error {
	_, err := redis.Set(context.Background(), key, objectId, time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}

func InitChunkUploadImur(key string) (string, error) {
	res, err := bucket.InitiateMultipartUpload(key)
	if err != nil {
		return "", err
	}
	return res.UploadID, nil
}

func GetChunkUploadImur(key string) (string, error) {
	objectId, err := redis.GetEx(context.Background(), key, time.Minute).Result()
	if err != nil {
		return "", err
	}
	return objectId, nil
}

func UploadFileChunk(chunk io.Reader, objectId string, size int64, index int) error {
	imur := oss.InitiateMultipartUploadResult{
		UploadID: objectId,
	}
	_, err := bucket.UploadPart(imur, chunk, size, index)
	return err
}
