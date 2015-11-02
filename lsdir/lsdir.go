package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := os.Args[1]
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
