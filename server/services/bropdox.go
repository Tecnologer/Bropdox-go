package services

import (
	"context"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/file"
	"github.com/tecnologer/bropdox/models/proto"
)

var (
	folderPath = "./files"
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

func (bs *BropdoxServer) Notifications(in *proto.File, stream proto.Bropdox_NotificationsServer) error {
	logrus.Debug("register for notifications")
	notifications := make(chan *proto.Response, 5)
	//dir, _ := os.Getwd()
	err := file.CreateWatcherRecursive(folderPath, notifications)
	if err != nil {
		logrus.WithError(err).Debug("error creating watcher")
		return err
	}

	logrus.Debug(folderPath)
	defer file.CloseWatchers()

	go func() {
		<-stream.Context().Done()
		close(notifications)
	}()

	for notif := range notifications {
		if reflect.TypeOf(notif.Content).String() == "*proto.Response_FileResponse" {
			(notif.Content.(*proto.Response_FileResponse)).FileResponse.File.Path = strings.Replace((notif.Content.(*proto.Response_FileResponse)).FileResponse.File.Path, folderPath, "", 1)
		}
		stream.Send(notif)
	}
	return nil
}
