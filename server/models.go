// Define some data models

package server

type FileMeta struct {
	FileId   string   `json:"file_id"`
	Name     string   `json:"name"`
	Suffix   string   `json:"suffix"`
	Size     int64    `json:"size"`
	UploadAt string   `json:"upload_at"`
	Tags     []string `json:"tags"`
}

type Tag struct {
	Name string `json:"name"`
	Desc string `json:"desciption"`
}
