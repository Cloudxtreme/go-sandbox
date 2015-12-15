package main

import (
	"log"
	"os"

	"github.com/go-fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool, 1)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	path, _ := os.Getwd()
	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
