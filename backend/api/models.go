// Define some data models

package api

type FileMeta struct {
	FileId   string `json:"file_id"`
	Name     string `json:"name"`
	Suffix   string `json:"suffix"`
	Size     int64  `json:"size"`
	UploadAt string `json:"upload_at"`
	Tags     []Tag  `json:"tags"` // Many-to-many relation
	Md5      string `json:"md5"`
}

type Color int

const (
	Red Color = iota
	Blue
	Green
	Yellow
	Purple
	Gray
)

func (c Color) ToString() string {
	colors := [...]string{"Red", "Blue", "Green", "Yellow", "Purple", "Gray"}
	return colors[c]
}

type Tag struct {
	Name  string `json:"name"`
	Desc  string `json:"desciption"`
	Color Color  `json:"color"`
}
