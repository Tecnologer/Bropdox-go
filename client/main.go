package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/client/src"
	"github.com/tecnologer/bropdox/models/proto"
)

var (
	verbouse    = flag.Bool("v", false, "enanble verbouse log")
	ip          = flag.String("ip", "", "ip of the server")
	port        = flag.Int("port", 8081, "port of the server")
	versionFlag = flag.Bool("version", false, "returns the version of the build")
	folderPath  = flag.String("path", "./files", "folder to store the files")

	version    string
	minversion string
	host       string
)

func init() {
	flag.Parse()
	if *verbouse {
		log.SetLevel(log.DebugLevel)
	}

	host = fmt.Sprintf("%s:%d", *ip, *port)
}

func main() {
	checkVersion()

	src.SetFolderPath(*folderPath)

	conn, err := src.GetConnectionAnonymous(host)
	if err != nil {
		log.Fatal(err)
	}
	defer src.CloseConnection()

	client := proto.NewBropdoxClient(conn)
	log.Infof("connected to %s", host)

	src.SyncFiles(client)
	go src.RegisterChangesRecursive(client, *folderPath)
	src.RegisterForNotifications(client)
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
	fmt.Printf("%s%s\n", version, minversion)
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
