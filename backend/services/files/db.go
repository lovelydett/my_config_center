package files

import (
	"wolf/drivers"
)

var db = drivers.GetPgConnector()

func insertFileMeta(filename string, tags []string, description string) error {
	params := map[string]interface{}{
		"filename":    filename,
		"tags":        tags,
		"description": description,
	}
	_, err := db.Conn.Exec("INSERT INTO files (filename, tags, description) VALUES (:filename, :tags, :description)", params)
	if err != nil {
		return err
	}
	return nil
}

func getImur(fileKey string) (string, error) {
	var uploadId string
	err := db.Conn.Get(&uploadId, "SELECT imur_upload_id FROM files WHERE md5=:key", map[string]interface{}{"key": fileKey})
	if err != nil {
		return "", err
	}
	return uploadId, nil
}

func getOrSetImur(fileKey string, uploadId string) (string, error) {
	txId := db.BeginTx()

}
