package convert

import (
	"os"
	"testing"
)

func TestCreateImageFile(t *testing.T) {
	_ = os.Remove("../testdata/to/icon_business_man01.jpg")

	convertErr := CreateImageFile("../testdata/from/icon_business_man01.png", "../testdata/to/icon_business_man01.jpg", "jpg")
	if convertErr != nil {
		t.Error("ファイルの変換に失敗しました")
	}

	if _, fileErr := os.Stat("../testdata/to/icon_business_man01.jpg"); os.IsNotExist(fileErr) {
		t.Error("変換後のファイルが見つかりません")
	}
	_ = os.Remove("../testdata/to/icon_business_man01.jpg")
}
