destDir=./dist/
binName=bropdox
formatVersion=+.%y%m%d
version=`git describe --tags`

both:
	make dir
	make server-local
	make client-local

server-local:
	go build -ldflags "-X main.minversion=`date -u $(formatVersion)` -X main.version=$(version)" -o "$(destDir)server/$(binName)" ./server/main.go

client-local:
	go build -ldflags "-X main.minversion=`date -u $(formatVersion)` -X main.version=$(version)" -o "$(destDir)client/$(binName)" ./client/main.go

dir:
	mkdir -p $(destDir)/client
	mkdir -p $(destDir)/server