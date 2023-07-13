package utils

import (
	"os"
)

func MkDirsIfNotExist(dirPath string) error {
	if !Exist(dirPath) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func Exist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}
