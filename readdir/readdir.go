package main

import (
	"log"
	"os"
)

func main() {

	dir, err := os.Open("/Users/zchee")
	// dir, err := os.Open(filepath.Join(os.Getenv("GOPATH"), "src/github.com"))
	checkErr(err)
	defer dir.Close()

	fi, err := dir.Stat()
	checkErr(err)

	filenames := make([]string, 0)

	if fi.IsDir() {
		fis, err := dir.Readdir(-1) // -1 means return all the FileInfos
		checkErr(err)

		for _, fileinfo := range fis {
			if !fileinfo.IsDir() {
				filenames = append(filenames, fileinfo.Name())
			}
		}
	}

	log.Println("Files: ", filenames)

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
