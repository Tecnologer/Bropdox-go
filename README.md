# Bropdox-Go

Remote File manager with Go + gRPC

## Code

1. `git clone https://github.com/Tecnologer/Bropdox-go.git $GOPATH/src/github.com/tecnologer/bropdox`
2. `cd $GOPATH/src/github.com/tecnologer/bropdox`
3. `git submodule init`
4. `git submodule update`
5. `cd server`
6. `go run main.go`

## Test With BloomRPC

1. Install [BloomRPC][1] and open it
2. Click "Import Protos"
3. Navigate to `$GOPATH/src/github.com/tecnologer/bropdox/proto`
4. Select `bropdox.proto`
5. Set address of the server, Default: `0.0.0.0:8081`
6. From the left panel, select the Message to test
7. Set the input for the message
8. Click "play"

### Example

![image](https://user-images.githubusercontent.com/8458967/114126065-493e6980-98bd-11eb-9a3a-6da121defed8.png)

## ToDo

- [x] Detect Changes
- [ ] Create File
- [ ] Update File
- [ ] Delete File
- [ ] gRPC Client
- [ ] GUI for Client

[1]: https://github.com/uw-labs/bloomrpc/releases