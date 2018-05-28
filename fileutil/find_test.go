package fileutil

import (
	"path/filepath"
	"testing"
)

func TestGetFileListSuccess(t *testing.T) {
	ret, _ := GetFileList("../testdata/from", "png")
	if len(ret) < 1 {
		t.Error("ファイルが見つかりません")
	}
	for _, path := range ret {
		if filepath.Ext(path) != ".png" {
			t.Error("誤ったファイルが見つかりました:" + path)
		}
	}
}

func TestGetFileListFailed(t *testing.T) {
	ret, _ := GetFileList("../testdata/from", "jpg")
	if len(ret) >= 1 {
		t.Error("見つからないはずのファイルが見つかりました")
	}
}
