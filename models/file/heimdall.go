package file

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/proto"
)

var watcher *fsnotify.Watcher

func CreateWatcher(path string, out chan<- *proto.Response) (err error) {
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return errors.Wrap(err, "creating watcher")
	}

	err = watcher.Add(path)
	if err != nil {
		return errors.Wrap(err, "adding path to watcher")
	}

	log.Debugf("register watcher at %s\n", path)
	go func() {
		for event := range watcher.Events {
			var eventType proto.TypeResponse

			isRemoved := event.Op&fsnotify.Remove == fsnotify.Remove

			if !isRemoved {
				isFolder, err := IsFolder(event.Name)
				if err != nil {
					out <- proto.ParseErrorToResponse(err)
					continue
				}

				if isFolder {
					continue
				}
			}

			switch {
			case event.Op&fsnotify.Write == fsnotify.Write:
				log.Debug("modified file:", event.Name)
				eventType = proto.TypeResponse_UPDATED
			case event.Op&fsnotify.Create == fsnotify.Create:
				log.Debug("created file:", event.Name)
				eventType = proto.TypeResponse_CREATED
			case isRemoved:
				log.Debug("deleted file:", event.Name)
				eventType = proto.TypeResponse_DELETED
			default:
				continue
			}

			var fileData *proto.File
			if isRemoved {
				fileData, _ = GetEmpty(event.Name)
			} else {
				fileData, err = Get(event.Name)
			}

			if err != nil {
				out <- proto.ParseErrorToResponse(err)
				continue
			}

			fileData.Path = strings.Replace(fileData.Path, path, "", 1)
			out <- proto.CreateFileResponse(fileData, eventType)
		}
	}()

	return nil
}

func CloseWatcher() error {
	if watcher != nil {
		watcher.Close()
	}
	return nil
}
