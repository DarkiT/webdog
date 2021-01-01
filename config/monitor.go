package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func InitMonitor(reloadFunc func()) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					reloadFunc()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("cfgMonitor error:", err)
			}
		}
	}()

	err = watcher.Add("./config.yml")
	if err != nil {
		log.Fatal("cfgMonitor error:", err)
	}
	<-done
}
