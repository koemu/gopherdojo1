package fileutil

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBatchConvert(t *testing.T) {
	var info ConvertInfo
	info.Filepath = []string{"../testdata/from/icon_business_man01.png"}
	info.FromExt = "png"
	info.ToDir = "../testdata/to"
	info.ToExt = "jpg"

	files, err := filepath.Glob(info.ToDir + "/*." + info.ToExt)
	if err != nil {
		t.Error(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			t.Error(err)
		}
	}

	convertErr := BatchConvert(&info)
	if convertErr != nil {
		t.Error(convertErr)
		t.Error("ファイルの変換に失敗しました")
	}

	convertedFiles, convertedErr := filepath.Glob(info.ToDir + "/*." + info.ToExt)
	if convertedErr != nil {
		t.Error(convertedErr)
	}
	if len(convertedFiles) < 1 {
		t.Error("変換後のファイルがありません")
	}
}
