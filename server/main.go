package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/proto"
	"github.com/tecnologer/bropdox/server/services"
	"google.golang.org/grpc"
)

var (
	verbouse    = flag.Bool("v", false, "enanble verbouse log")
	port        = flag.Int("port", 8081, "port of host to expose the server")
	versionFlag = flag.Bool("version", false, "returns the version of the build")

	version    string
	minversion string
	host       string
)

func init() {
	flag.Parse()
	if *verbouse {
		log.SetLevel(log.DebugLevel)
	}

	host = fmt.Sprintf(":%d", *port)
}

func main() {
	checkVersion()
	listener, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatalf("Unable to listen '%s': %v\n", host, err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	srv := &services.BropdoxServer{}

	proto.RegisterBropdoxServer(s, srv)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}
	}()

	log.Infof("Server started at host %s\n", host)
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	// fmt.Println("Closing MongoDB connection")
	// db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}

func checkVersion() {
	if len(os.Args) < 2 {
		return
	}

	if *versionFlag || argsConstainsVersion() {
		printVersion()
	}
}

func printVersion() {
	fmt.Printf("bropdox server %s%s\n", version, minversion)
	os.Exit(0)
}

func argsConstainsVersion() bool {
	for _, a := range os.Args {
		if a == "version" || a == "--version" {
			return true
		}
	}

	return false
}
