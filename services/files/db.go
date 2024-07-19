package files

import (
	"wolf/drivers"
)

var mysql = drivers.GetMySQLConnector()

func insertFileMeta(filename string, tags []string, description string) error {
	_, err := mysql.Exec("INSERT INTO files (filename, tags, description) VALUES (?, ?, ?)", filename, tags, description)
	if err != nil {
		return err
	}
	return nil
}
