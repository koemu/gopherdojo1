//
package fileutil

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/koemu/gopherdojo1/convert"
)

// ConvertInfo 変換情報
type ConvertInfo struct {
	Filepath []string
	FromExt  string
	ToDir    string
	ToExt    string
}

// BatchConvert 一括変換
func BatchConvert(info *ConvertInfo) error {
	for _, path := range info.Filepath {
		filename := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
		toPath := filepath.Join(info.ToDir, filename+"."+info.ToExt)
		if _, err := os.Stat(toPath); !os.IsNotExist(err) {
			return errors.New("File exists:" + toPath)
		}
		err := convert.CreateImageFile(path, toPath, info.ToExt)
		if err != nil {
			return err
		}
	}

	return nil
}
