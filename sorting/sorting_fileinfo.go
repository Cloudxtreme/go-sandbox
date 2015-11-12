// Base of http://qiita.com/shinofara/items/e5e78e6864a60dc851a6
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
)

type byName []os.FileInfo

func (f byName) Len() int           { return len(f) }
func (f byName) Less(i, j int) bool { return f[i].Name() < f[j].Name() }
func (f byName) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func main() {
	var cDir, _ = os.Getwd()
	cDir += "/"

	var _, filePattern = path.Split(cDir)

	fileInfos, err := ioutil.ReadDir(cDir)
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(byName(fileInfos))
	for _, fileInfo := range fileInfos {
		var findName = (fileInfo).Name()
		var findTime = (fileInfo).ModTime().Unix()

		var matched = true
		if filePattern != "" {
			matched, _ = path.Match(filePattern, findName)
		}
		if matched == true {
			fmt.Printf("%s\n", findName)
			fmt.Println(findTime)
		}
	}
}
