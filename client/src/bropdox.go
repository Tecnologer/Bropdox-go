package src

import (
	"context"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/files"
	"github.com/tecnologer/bropdox/models/proto"
)

var (
	lastFilesUpdates = make(map[string]proto.TypeResponse)
)

func SyncFiles(client proto.BropdoxClient) error {
	response, err := client.GetFiles(context.Background(), &proto.Empty{})

	if err != nil {
		log.WithError(err).Warn("getting list of files")
	}

	filesResponse := response.GetFilesResponse()

	log.Info("Existing files on server: ")

	localFiles, err := files.GetListFileRecursive(folderPath)
	if err != nil {
		log.Warnf("sync files: getting list of local files")

	}

	if filesResponse != nil {
		remoteFiles := []string{}
		for _, file := range filesResponse.Files {
			fmt.Printf("\t- %s\n", file.Path)

			remoteFiles = append(remoteFiles, file.Path)

			fileRes, err := client.GetFile(context.Background(), file)
			if err != nil {
				log.WithError(err).Warnf("sync files: %s")
				continue
			}
			eRes := fileRes.GetErrorResponse()
			if eRes != nil {
				log.Warnf("sync files: trying get file. Response: %s. Path: %s", eRes.Message, file.Path)
				continue
			}

			fRes := fileRes.GetFileResponse()

			path := fmt.Sprintf("%s%s", folderPath, fRes.File.Path)
			path = strings.Replace(path, "//", "/", -1)
			err = files.CreateOrUpdate(path, fRes.File.Content)
			if err != nil {
				log.WithError(err).Warnf("sync files: downloading file %s", path)
				continue
			}
		}

		compareLocalFilesWRemote(localFiles, remoteFiles)
	}
	return nil
}

func compareLocalFilesWRemote(localFiles, remoteFiles []string) {
	exists := false
	length := utf8.RuneCountInString(folderPath)
	for _, local := range localFiles {
		exists = false
		local = local[length-2:]
		for _, remote := range remoteFiles {
			if local == remote {
				exists = true
				break
			}
		}

		if !exists {
			path := fmt.Sprintf("%s%s", folderPath, local)
			path = strings.Replace(path, "//", "/", -1)
			files.Remove(path)
		}
	}
}

func RegisterForNotifications(client proto.BropdoxClient) error {
	notifReq := &proto.NotificationsRequest{
		Id: fmt.Sprint(time.Now().Unix()),
	}

	notifications, err := client.Notifications(context.Background(), notifReq)
	if err != nil {
		log.WithError(err).Warn("register for notifications")
	}

	for {
		response, err := notifications.Recv()
		if err != nil {
			log.WithError(err).Warn("new error notification")
			return errors.Wrap(err, "error received instead of notification")
		}
		errorResponse := response.GetErrorResponse()
		if errorResponse != nil {
			log.Warnf("notification of error: %s", errorResponse.Message)
			continue
		}

		fileResponse := response.GetFileResponse()
		log.WithFields(log.Fields{
			"path": fileResponse.File.Path,
			"type": fileResponse.Type,
		}).Debug("new notification")

		f := fileResponse.File
		lastFilesUpdates[f.Path] = fileResponse.Type
		path := fmt.Sprintf("%s%s", folderPath, f.Path)

		switch fileResponse.Type {
		case proto.TypeResponse_UPDATED:
			fallthrough
		case proto.TypeResponse_CREATED:
			err = files.CreateOrUpdate(path, f.Content)
			if err != nil {
				log.WithError(err).Errorf("file notif: %s", fileResponse.Type)
			}
		case proto.TypeResponse_DELETED:
			err = files.Remove(path)
			if err != nil {
				log.WithError(err).Errorf("removing file %s", path)
			}
		}
	}
}
