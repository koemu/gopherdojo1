package fileutil

import (
	"os"
	"path/filepath"
)

// GetFileList retrieve target file list
func GetFileList(frompath string, ext string) ([]string, error) {
	var filelist []string

	err := filepath.Walk(frompath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == "."+ext {
			filelist = append(filelist, path)
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return filelist, nil
}
