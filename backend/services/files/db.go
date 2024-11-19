package files

import (
	"context"
	"time"
	"wolf/drivers"
)

var pg = drivers.Pg
var redis = drivers.Redis

func insertFileMeta(filename string, tags []string, description string) error {
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

func initFileUpload(key string) bool {
	val, _ := redis.SetNX(context.Background(), key, "inited", 10*time.Minute).Result()
	if val {
		return true
	}
	return false
}

func completeFileUploadInit(key string) error {
	_, err := redis.Set(context.Background(), key, "wip", 10*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}
