package services

import (
	"errors"
	"io"
	"net/http"
	"time"

	"wolf/db"
	"wolf/utils"
)

type Chunk struct {
	key   string
	index int
	total int
	size  int
	data  io.Reader
}

func ParseChunk(r *http.Request) (Chunk, error) {
	r.ParseMultipartForm(10 << 20) // 10M
	data, _, err := r.FormFile("chunk")
	if err != nil {
		return Chunk{}, err
	}
	return Chunk{
		key:   r.FormValue("key"),
		index: utils.Atoi(r.FormValue("index")),
		total: utils.Atoi(r.FormValue("total")),
		size:  utils.Atoi(r.FormValue("size")),
		data:  data,
	}, nil
}

func UploadChunk(chunk Chunk) error {
	var objectId string
	var err error

	// Step1: Try the distributed lock
	if db.InitChunkUpload(chunk.key) {
		// Got the lock, responsible for registering the task
		objectId, err = db.InitChunkUploadImur(chunk.key)
		if err != nil {
			return err
		}
		// Update the task status
		err = db.UpdateChunkUploadImur(chunk.key, objectId)
		if err != nil {
			return err
		}
	}

	// By default, the task is already registered
	retry := 0
	for retry < 3 {
		// Try to get the task status
		objectId, err = db.GetChunkUploadImur(chunk.key)
		if err != nil {
			return err
		}
		if objectId != "" && objectId != "inited" {
			break
		}
		// Sleep for a while
		time.Sleep(1 * time.Second)
		retry += 1
	}

	if retry >= 3 {
		return errors.New("max retry exceeded for upload ID check")
	}

	// Step2: Upload the file chunk
	err = db.UploadFileChunk(chunk.data, objectId, int64(chunk.size), chunk.index)
	if err != nil {
		return err
	}

	return nil
}
