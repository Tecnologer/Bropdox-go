# Bropdox-Go

Remote File manager with Go + gRPC

## How to use it?

### Server

1. [Download][2] your version
1. Run the binary
1. Use the folling flags
   ```txt
   -port int
      port of host to expose the server (default 8081)
   -v	enanble verbouse log
   -version
    	returns the version of the build
   ```

Example:

```bash
./bropdox-server_<version_os_arch> -port 8080
```

### Client

1. [Download][2] your version
1. Run the binary
1. Use the folling flags

   ```txt
     -ip string
        ip of the server
    -path string
        folder to store the files (default "./files")
    -port int
        port of the server (default 8081)
    -v	enanble verbouse log
    -version
        returns the version of the build

   ```

Example:

```bash
./bropdox-server_<version_os_arch> -port 8080 -path '~/Documents/bropdox-files'
```

## Code

1. `git clone https://github.com/Tecnologer/Bropdox-go.git $GOPATH/src/github.com/tecnologer/bropdox`
2. `cd $GOPATH/src/github.com/tecnologer/bropdox`
3. `git submodule init`
4. `git submodule update`
5. `cd server`, change to `cd client` to run the client version
6. `go run main.go`

### Generate binaries with Make

You could use **Make** to generate binary files for client and server:

- For client:
  `make client-local`
- For server:
  `make server-local`

> The binaries will be stored in `./dist/<server|client>`

Note: You can use onle `make` to generate both binaries.

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

- Server Side
  - [x] Detect Changes
  - [x] Create File
  - [x] Update File
  - [x] Delete File
  - [x] Get File
  - [x] Get Files (list)
  - [ ] Support multiple clients
  - [ ] Use flag to set the root folder
- Client Side
  - [x] Detect Changes
  - [x] Send File Created
  - [x] Send File Updated
  - [x] File Deleted
  - [ ] Check status on startup and sync
    - [x] Delete files deleted on server
    - [ ] Update files created offline
  - [x] Register for notifications
  - [x] Update files on local from notifications (Create, Update & Remove)
- Makefile
  - [ ] Support cross-compilation

[1]: https://github.com/uw-labs/bloomrpc/releases
[2]: https://github.com/Tecnologer/Bropdox-go/releases
