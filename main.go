package main

import (
	"flag"
	"fmt"

	"github.com/koemu/gopherdojo1/fileutil"
)

/**
 * 次の仕様を満たすコマンドを作って下さい
 * ディレクトリを指定する
 * 指定したディレクトリ以下のJPGファイルをPNGに変換
 * ディレクトリ以下は再帰的に処理する
 * 変換前と変換後の画像形式を指定できる
 * 以下を満たすように開発してください
 * mainパッケージと分離する
 * 自作パッケージと標準パッケージと準標準パッケージのみ使う
 * 準標準パッケージ：golang.org/x以下のパッケージ
 * ユーザ定義型を作ってみる
 * GoDocを生成してみる
 */

func main() {
	from := flag.String("from", "/tmp/from", "インポート元")
	to := flag.String("to", "/tmp/to", "インポート先")
	fromFormat := flag.String("from_format", "png", "インポート元フォーマット")
	toFormat := flag.String("to_format", "jpg", "インポート先フォーマット")
	flag.Parse()

	filelist, err := fileutil.GetFileList(*from, *fromFormat)
	if err != nil {
		panic(err)
	}

	var info fileutil.ConvertInfo
	info.Filepath = filelist
	info.FromExt = *fromFormat
	info.ToDir = *to
	info.ToExt = *toFormat
	err = fileutil.BatchConvert(&info)

	if err != nil {
		panic(err)
	}

	fmt.Println("Convert have been finished.")
}
