package main

import (
	"net/http"
	"os"
)

func main() {
	path, _ := os.Getwd()
	http.Handle("/", http.FileServer(http.Dir(path)))
	http.ListenAndServe(":3000", nil)
}
