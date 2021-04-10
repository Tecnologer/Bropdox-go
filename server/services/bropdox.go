package services

import (
	"context"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/file"
	"github.com/tecnologer/bropdox/models/proto"
)

type BropdoxServer struct{}

func (bs *BropdoxServer) CreateFile(ctx context.Context, in *proto.File) (*proto.Response, error) {
	return proto.CreateFileResponse(in, proto.TypeResponse_CREATED), nil
}

func (bs *BropdoxServer) UpdateFile(ctx context.Context, in *proto.File) (*proto.Response, error) {
	return proto.CreateFileResponse(in, proto.TypeResponse_UPDATED), nil
}

func (bs *BropdoxServer) RemoveFile(ctx context.Context, in *proto.File) (*proto.Response, error) {
	return proto.CreateFileResponse(in, proto.TypeResponse_DELETED), nil
}

func (bs *BropdoxServer) GetFile(ctx context.Context, in *proto.File) (*proto.Response, error) {
	return proto.CreateFileResponse(in, proto.TypeResponse_UPDATED), nil
}

func (bs *BropdoxServer) GetFiles(ctx context.Context, _ *proto.Empty) (*proto.Response, error) {
	return proto.CreateFilesResponse([]*proto.File{}), nil
}

func (bs *BropdoxServer) Notifications(in *proto.File, stream proto.Bropdox_NotificationsServer) error {
	logrus.Debug("register for notifications")
	notifications := make(chan *proto.Response, 5)
	//dir, _ := os.Getwd()
	path := "./files"
	err := file.CreateWatcherRecursive(path, notifications)
	if err != nil {
		logrus.WithError(err).Debug("error creating watcher")
		return err
	}

	logrus.Debug(path)
	defer file.CloseWatchers()

	go func() {
		<-stream.Context().Done()
		close(notifications)
	}()

	for notif := range notifications {
		if reflect.TypeOf(notif.Content).String() == "*proto.Response_FileResponse" {
			(notif.Content.(*proto.Response_FileResponse)).FileResponse.File.Path = strings.Replace((notif.Content.(*proto.Response_FileResponse)).FileResponse.File.Path, path, "", 1)
		}
		stream.Send(notif)

	}
	return nil
}
