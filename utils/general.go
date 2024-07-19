package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// GetProjectRoot returns the root directory of the project.
var projectRoot string

func GetProjectRoot() string {
	if projectRoot != "" {
		return projectRoot
	}
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			projectRoot = dir
			return dir
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}

		dir = parentDir
	}

	return ""
}
