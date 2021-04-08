package services

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/proto"
	"github.com/tecnologer/bropdox/server/file"
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

func (bs *BropdoxServer) Notifications(in *proto.File, stream proto.Bropdox_NotificationsServer) error {
	logrus.Debug("register for notifications")
	notifications := make(chan *proto.Response, 5)
	dir, _ := os.Getwd()
	path := dir + "/files"
	err := file.CreateWatcher(path, notifications)
	if err != nil {
		logrus.WithError(err).Debug("error creating watcher")
		return err
	}

	logrus.Debug(path)
	defer file.CloseWatcher()

	for notif := range notifications {
		stream.Send(notif)
	}
	return nil
}

func (bs *BropdoxServer) mustEmbedUnimplementedBropdoxServer() {}
