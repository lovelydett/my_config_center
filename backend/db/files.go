package db

import (
	"context"
	"io"
	"time"
	"wolf/config"
	"wolf/drivers"

	"github.com/redis/go-redis/v9"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var pg = drivers.Pg
var rd = drivers.Redis
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
	val, _ := rd.SetNX(context.Background(), key, "inited", time.Minute).Result()
	return val
}

func UpdateChunkUploadImur(key string, uploadId string) error {
	_, err := rd.Set(context.Background(), key, uploadId, time.Minute).Result()
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
	uploadId, err := rd.GetEx(context.Background(), key, time.Minute).Result()
	if err != nil {
		return "", err
	}
	return uploadId, nil
}

func UploadFileChunk(chunk io.Reader, key string, uploadId string, size int64, index int) error {
	imur := oss.InitiateMultipartUploadResult{
		Bucket:   config.GetDeployConfig().OSS.BucketName,
		Key:      key,
		UploadID: uploadId,
	}
	part, err := bucket.UploadPart(imur, chunk, size, index)
	if err != nil {
		return err
	}

	zsetKey := "zset_" + key

	_, err = rd.ZAdd(context.TODO(), zsetKey, redis.Z{
		Score:  float64(part.PartNumber),
		Member: part.ETag,
	}).Result()

	if err != nil {
		return err
	}

	_, err = rd.Expire(context.TODO(), zsetKey, time.Minute).Result()

	return err
}

func GetUploadedParts(key string) ([]oss.UploadPart, error) {
	zsetKey := "zset_" + key
	res, err := rd.ZRangeWithScores(context.TODO(), zsetKey, 0, -1).Result()

	if err != nil {
		return nil, err
	}

	parts := make([]oss.UploadPart, 0)

	for _, each := range res {
		parts = append(parts, oss.UploadPart{
			ETag:       each.Member.(string),
			PartNumber: int(each.Score),
		})
	}

	return parts, nil
}

func CompleteChunkUploadToOSS(uploadId string, key string, parts []oss.UploadPart) error {
	imur := oss.InitiateMultipartUploadResult{
		Bucket:   config.GetDeployConfig().OSS.BucketName,
		Key:      key,
		UploadID: uploadId,
	}

	objectAcl := oss.ObjectACL(oss.ACLPrivate)
	_, err := bucket.CompleteMultipartUpload(imur, parts, objectAcl)

	return err
}
