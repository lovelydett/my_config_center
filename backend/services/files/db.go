package files

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

func insertChunkMeta(filename string, tags []string, description string) error {
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

func initChunkUpload(key string) bool {
	val, _ := redis.SetNX(context.Background(), key, "inited", time.Minute).Result()
	if val {
		return true
	}
	return false
}

func updateChunkUploadImur(key string, objectId string) error {
	_, err := redis.Set(context.Background(), key, objectId, time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}

func initChunkUploadImur(key string) (string, error) {
	res, err := bucket.InitiateMultipartUpload(key)
	if err != nil {
		return "", err
	}
	return res.UploadID, nil
}

func getChunkUploadImur(key string) (string, error) {
	objectId, err := redis.GetEx(context.Background(), key, time.Minute).Result()
	if err != nil {
		return "", err
	}
	return objectId, nil
}

func uploadFileChunk(chunk io.Reader, objectId string, size int64, index int) error {
	imur := oss.InitiateMultipartUploadResult{
		UploadID: objectId,
	}
	_, err := bucket.UploadPart(imur, chunk, size, index)
	return err
}
