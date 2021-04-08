package services

import (
	"context"

	"github.com/tecnologer/bropdox/models/proto"
)

type BropdoxServer struct{}

func (bs *BropdoxServer) CreateFile(ctx context.Context, in *proto.File) (*proto.FileResponse, error) {
	return &proto.FileResponse{
		File: in,
		Type: proto.TypeResponse_CREATED,
	}, nil
}

func (bs *BropdoxServer) UpdateFile(ctx context.Context, in *proto.File) (*proto.FileResponse, error) {
	return &proto.FileResponse{
		File: in,
		Type: proto.TypeResponse_UPDATED,
	}, nil
}

func (bs *BropdoxServer) RemoveFile(ctx context.Context, in *proto.File) (*proto.FileResponse, error) {
	return &proto.FileResponse{
		File: in,
		Type: proto.TypeResponse_DELETED,
	}, nil
}

func (bs *BropdoxServer) Notifications(*proto.File, proto.Bropdox_NotificationsServer) error {
	return nil
}

func (bs *BropdoxServer) mustEmbedUnimplementedBropdoxServer() {}
