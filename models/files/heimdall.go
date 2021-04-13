package files

import (
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/proto"
)

var (
	watchers      = NewWatcherCollection()
	isRecursively bool
)

func CreateWatcherRecursive(path string, out chan<- *proto.Response) (err error) {
	isRecursively = true

	err = addWatcherRecursive(path, out)
	if err != nil {
		return errors.Wrapf(err, "creating watcher recursively for %s", path)
	}

	return nil
}

func addWatcherRecursive(path string, out chan<- *proto.Response) (err error) {
	watcher, err := CreateWatcher(path, out)
	if err != nil {
		return errors.Wrapf(err, "adding new watcher for path %s", path)
	}
	watchers.Add(path, watcher)

	subFolders, err := listFolders(path)
	if err != nil {
		return errors.Wrapf(err, "creating watcher for subfolders of %s", path)
	}

	var folderPath string
	for _, folder := range subFolders {
		folderPath = path + "/" + folder

		err := addWatcherRecursive(folderPath, out)
		if err != nil {
			return errors.Wrapf(err, "getting subfolders recursively for %s", folder)
		}
	}

	return nil
}

func CreateWatcher(path string, out chan<- *proto.Response) (watcher *fsnotify.Watcher, err error) {
	watcher, exists := watchers.get(path)
	if exists {
		return watcher, nil
	}

	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return nil, errors.Wrap(err, "creating watcher")
	}

	err = watcher.Add(path)
	if err != nil {
		return nil, errors.Wrap(err, "adding path to watcher")
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
					if isRecursively {
						err = addWatcherRecursive(event.Name, out)
						if err != nil {
							out <- proto.ParseErrorToResponse(err)
						}
					}
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
			if eventType == proto.TypeResponse_UPDATED {
				fileData, err = Get(event.Name)
			} else {
				fileData, _ = GetEmpty(event.Name)
			}

			if err != nil {
				out <- proto.ParseErrorToResponse(err)
				continue
			}

			out <- proto.CreateFileResponse(fileData, eventType)
		}
	}()

	return
}

func CloseWatchers() error {
	watchers.Clear()
	return nil
}
