package src

import (
	"context"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/files"
	"github.com/tecnologer/bropdox/models/proto"
)

var (
	folderPath string
)

func SetFolderPath(path string) {
	folderPath = path

	err := files.MkdirIfNotExists(folderPath)
	if err != nil {
		log.WithError(err).Errorf("creating the root folder: %s", folderPath)
	}
}

func RegisterChangesRecursive(client proto.BropdoxClient, path string) error {
	notifications := make(chan *proto.Response, 5)

	err := files.CreateWatcherRecursive(path, notifications)
	if err != nil {
		log.WithError(err).Debug("error creating watcher")
		return err
	}

	for notif := range notifications {
		fileRes := notif.GetFileResponse()

		if fileRes == nil {
			continue
		}
		file := fileRes.File

		//fileLocalPath := fmt.Sprintf("%s%s", folderPath, file.Path)
		file.Path = strings.Replace(file.Path, path, "", 1)

		if !fileWasUpdatedLocally(file.Path, fileRes.Type) {
			continue
		}

		switch fileRes.Type {
		case proto.TypeResponse_CREATED:
			client.CreateFile(context.Background(), file)
		case proto.TypeResponse_UPDATED:
			client.UpdateFile(context.Background(), file)
		case proto.TypeResponse_DELETED:
			client.RemoveFile(context.Background(), file)
		}
	}
	return nil
}

func fileWasUpdatedLocally(path string, t proto.TypeResponse) bool {
	lastType, exists := lastFilesUpdates[path]
	if !exists {
		return true
	}

	return lastType != t
}
