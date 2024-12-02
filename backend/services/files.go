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
	// Currently only support upload to OSS
	return uploadChunkToOSS(chunk)
}

func MergeChunks(key string) error {
	// Currently only support merge the file chunks in OSS
	return mergeChunksInOSS(key)
}

func uploadChunkToOSS(chunk Chunk) error {
	var uploadId string
	var err error

	// Step1: Try the distributed lock
	if db.InitChunkUpload(chunk.key) {
		// Got the lock, responsible for registering the task
		uploadId, err = db.InitChunkUploadImur(chunk.key)
		if err != nil {
			return err
		}
		// Update the task status
		err = db.UpdateChunkUploadImur(chunk.key, uploadId)
		if err != nil {
			return err
		}
	}

	// By default, the task is already registered
	retry := 0
	for retry < 3 {
		// Try to get the task status
		uploadId, err = db.GetChunkUploadImur(chunk.key)
		if err != nil {
			return err
		}
		if uploadId != "" && uploadId != "inited" {
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
	err = db.UploadFileChunk(chunk.data, uploadId, uploadId, int64(chunk.size), chunk.index)
	if err != nil {
		return err
	}

	return nil
}

func mergeChunksInOSS(key string) error {
	uploadId, err := db.GetChunkUploadImur(key)
	if err != nil {
		return err
	}

	return db.CompleteChunkUploadToOSS(uploadId, key)
}
