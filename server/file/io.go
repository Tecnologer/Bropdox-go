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

func Get(path string) (*proto.File, error) {
	open, e := os.Open(path)
	if e != nil {
		return nil, errors.Wrap(e, "file get: opening")
	}

	data, e := io.ReadAll(open)
	if e != nil {
		return nil, errors.Wrap(e, "file get: reading")
	}

	file := &proto.File{
		Path:    path,
		Content: data,
	}

	return file, nil
}
