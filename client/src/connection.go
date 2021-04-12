package src

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var (
	connection *grpc.ClientConn
)

func GetConnection(host string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	var err error
	CloseConnection()

	connection, err = grpc.Dial(host, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "connecting to %s", host)
	}

	return connection, nil
}

func CloseConnection() {
	if connection != nil {
		connection.Close()
	}
}

func GetConnectionAnonymous(host string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	_, err := GetConnection(host, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "connecting with no credentials to %s", host)
	}

	return connection, nil
}
