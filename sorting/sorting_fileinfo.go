// Base of http://qiita.com/shinofara/items/e5e78e6864a60dc851a6
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
)

type FileInfos []os.FileInfo
type ByName struct{ FileInfos }

func (fi ByName) Len() int {
	return len(fi.FileInfos)
}
func (fi ByName) Swap(i, j int) {
	fi.FileInfos[i], fi.FileInfos[j] = fi.FileInfos[j], fi.FileInfos[i]
}
func (fi ByName) Less(i, j int) bool {
	return fi.FileInfos[j].ModTime().Unix() < fi.FileInfos[i].ModTime().Unix()
}

// 指定されたファイル名がディレクトリかどうか調べる
func IsDirectory(name string) (isDir bool, err error) {
	fInfo, err := os.Stat(name) // FileInfo型が返る。
	if err != nil {
		return false, err // もしエラーならエラー情報を返す
	}
	// ディレクトリかどうかチェック
	return fInfo.IsDir(), nil
}

func main() {
	var arg string

	// -fオプション flag.Arg(0)だとファイル名が展開されてしまうようなので
	flag.StringVar(&arg, "f", "", "SearchPattern")
	// コマンドライン引数を解析
	flag.Parse()

	// カレントディレクトリの取得
	var curDir, _ = os.Getwd()
	curDir += "/"

	// 引数が取得できなければ、カレントディレクトリを使用
	if arg == "" {
		arg = curDir
	}

	// ディレクトリとファイル名に分割して格納
	var dirName, filePattern = path.Split(arg)

	// ディレクトリが無いならばカレントディレクトリを使用
	if dirName == "" {
		dirName = curDir
	}

	// 取得しようとしているパスがディレクトリかチェック
	var isDir, _ = IsDirectory(dirName + filePattern)

	// ディレクトリならば、そのディレクトリ配下のファイルを調べる。
	if isDir == true {
		dirName = dirName + filePattern
		filePattern = ""
	}

	// ディレクトリ内のファイル情報の読み込み[] *os.FileInfoが返る。
	fileInfos, err := ioutil.ReadDir(dirName)

	// ディレクトリの読み込みに失敗したらエラーで終了
	if err != nil {
		fmt.Errorf("Directory cannot read %s\n", err)
		os.Exit(1)
	}

	// ファイル情報を一つずつ表示する
	sort.Sort(ByName{fileInfos})
	for _, fileInfo := range fileInfos {
		// *FileInfo型
		var findName = (fileInfo).Name()
		var matched = true
		// lsのようなワイルドカード検索を行うため、path.Matchを呼び出す
		if filePattern != "" {
			matched, _ = path.Match(filePattern, findName)
		}
		// path.Matchでマッチした場合、ファイル名を表示
		if matched == true {
			fmt.Printf("%s\n", findName)
		}
	}
}
