package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/proto"
	"google.golang.org/grpc"
)

var (
	verbouse = flag.Bool("v", false, "enanble verbouse log")
	ip       = flag.String("ip", "", "ip of the server")
	port     = flag.Int("port", 8081, "port of the server")

	host string
)

func init() {
	flag.Parse()
	if *verbouse {
		log.SetLevel(log.DebugLevel)
	}

	host = fmt.Sprintf("%s:%d", *ip, *port)
}

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewBropdoxClient(conn)
	log.Infof("connected to %s", host)

	response, err := client.GetFiles(context.Background(), &proto.Empty{})

	if err != nil {
		log.WithError(err).Warn("getting list of files")
	}

	filesResponse := response.GetFilesResponse()

	log.Info("Existing files on server: ")
	for _, file := range filesResponse.Files {
		fmt.Printf("\t- %s\n", file.Path)
	}
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
			log.WithError(err).Warn("register for notifications")
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
		}).Info("new notification")
	}
	log.Infof("started at %s", host)
}
