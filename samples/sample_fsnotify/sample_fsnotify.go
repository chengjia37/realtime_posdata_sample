package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run " + filepath.Base(os.Args[0]) + ".go [directory_name]")
		return
	}
	dirname := os.Args[1]

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go procesDirChange(watcher, done)

	err = watcher.Add(dirname)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

func procesDirChange(watcher *fsnotify.Watcher, done chan bool) {
	for {
		select {
		case event := <-watcher.Events:
			fmt.Println("events: ", event)
			switch {
			case event.Op&fsnotify.Create == fsnotify.Create:
				log.Println("Created file: ", event.Name)
			}
		case err := <-watcher.Errors:
			log.Println("error: ", err)
			done <- true
		}
	}
}
