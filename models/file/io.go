package file

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/tecnologer/bropdox/models/proto"
)

func Create(path string, content []byte) error {
	return nil
}

func GetContent(path string) ([]byte, error) {
	return nil, nil
}

func GetEmpty(path string) (*proto.File, error) {
	return &proto.File{
		Path: path,
	}, nil
}

func Get(path string) (*proto.File, error) {
	open, e := os.Open(path)
	if e != nil {
		return nil, errors.Wrap(e, "file get: opening")
	}

	data, e := io.ReadAll(open)
	if e != nil {
		return nil, errors.Wrap(e, "file get: reading")
	}

	file, _ := GetEmpty(path)
	file.Content = data

	return file, nil
}

func IsFolder(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, errors.Wrapf(err, "isFolder: getting stats for %s", path)
	}

	return fi.Mode().IsDir(), nil
}
