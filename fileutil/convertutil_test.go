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
		t.Error("ファイルの変換に失敗しました")
	}

	if _, fileErr := os.Stat("../testdata/to/icon_business_man01.jpg"); os.IsNotExist(fileErr) {
		t.Error("変換後のファイルが見つかりません")
	}
	_ = os.Remove("../testdata/to/icon_business_man01.jpg")
}
