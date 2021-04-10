package services

import (
	"context"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/file"
	"github.com/tecnologer/bropdox/models/proto"
)

var (
	folderPath = "./files"
	clients    = map[string]proto.Bropdox_NotificationsServer{}
)

type BropdoxServer struct{}

func (bs *BropdoxServer) CreateFile(ctx context.Context, in *proto.File) (*proto.Response, error) {
	err := file.Create(folderPath+"/"+in.Path, in.Content)
	if err != nil {
		return proto.ParseErrorToResponse(err), nil
	}

	return proto.CreateFileResponse(in, proto.TypeResponse_CREATED), nil
}

func (bs *BropdoxServer) UpdateFile(ctx context.Context, in *proto.File) (*proto.Response, error) {
	err := file.Create(folderPath+"/"+in.Path, in.Content)
	if err != nil {
		return proto.ParseErrorToResponse(err), nil
	}
	return proto.CreateFileResponse(in, proto.TypeResponse_UPDATED), nil
}

func (bs *BropdoxServer) RemoveFile(ctx context.Context, in *proto.File) (*proto.Response, error) {
	err := file.Remove(folderPath + "/" + in.Path)
	if err != nil {
		return proto.ParseErrorToResponse(err), nil
	}
	return proto.CreateFileResponse(in, proto.TypeResponse_DELETED), nil
}

func (bs *BropdoxServer) GetFile(ctx context.Context, in *proto.File) (*proto.Response, error) {
	file, err := file.Get(folderPath + "/" + in.Path)
	if err != nil {
		return proto.ParseErrorToResponse(err), nil
	}

	file.Path = strings.Replace(file.Path, folderPath, "", 1)
	return proto.CreateFileResponse(file, proto.TypeResponse_UPDATED), nil
}

func (bs *BropdoxServer) GetFiles(ctx context.Context, _ *proto.Empty) (*proto.Response, error) {
	files, err := file.GetListFileRecursive(folderPath)
	if err != nil {
		return proto.ParseErrorToResponse(err), nil
	}

	filesRes := make([]*proto.File, 0)
	for _, file := range files {
		if file == folderPath {
			continue
		} else if !strings.HasPrefix(file, "./") {
			file = "./" + file
		}
		file = strings.Replace(file, folderPath, "", 1)
		fileRes := proto.NewFile(file)

		filesRes = append(filesRes, fileRes)
	}

	return proto.CreateFilesResponse(filesRes), nil
}

func (bs *BropdoxServer) Notifications(in *proto.NotificationsRequest, stream proto.Bropdox_NotificationsServer) error {
	log.Debug("register for notifications")
	if in.Id == "" {
		return fmt.Errorf("the client ID is required")
	}

	notifications := make(chan *proto.Response, 5)
	//dir, _ := os.Getwd()
	err := file.CreateWatcherRecursive(folderPath, notifications)
	if err != nil {
		log.WithError(err).Debug("error creating watcher")
		return err
	}

	log.Debugf("new client: %s", in.Id)
	clients[in.Id] = stream

	defer file.CloseWatchers()

	go func() {
		<-stream.Context().Done()
		close(notifications)
	}()

	for notif := range notifications {
		fileRes := notif.GetFileResponse()
		if fileRes != nil {
			fileRes.File.Path = strings.Replace(fileRes.File.Path, folderPath, "", 1)
		}

		for _, stm := range clients {
			stm.Send(notif)
		}
	}
	return nil
}
