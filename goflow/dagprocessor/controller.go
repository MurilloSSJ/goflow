package dagprocessor

import (
	"lessflow/v2/core/handlers/config"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func RunProcessor() {
    config := config.LoadConfig()
    dagFolder := config.Base.DagDir
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
		log.Fatalf("Error on create watcher: %v", err)
        panic(err)
	}
    defer watcher.Close()
    err = watcher.Add(dagFolder)
    if err != nil {
        log.Fatalf("Error on watch dag folder: %v", err)
        panic(err)
    }
    done := make(chan bool)
    go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				dagProcessor := Constructor(event.Name)
				fileSplited := strings.Split(event.Name, ".")
				if !ok || fileSplited[len(fileSplited)-1] != "go" {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Printf("Arquivo criado: %s", event.Name)
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Printf("Arquivo modificado: %s", event.Name)
					dagProcessor.UpdateDag()
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Printf("Arquivo removido: %s", event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Printf("Arquivo renomeado: %s", event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Erro: %v", err)
			}
		}
	}()
	<-done

}