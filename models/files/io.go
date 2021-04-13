package files

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/bropdox/models/proto"
)

func CreateOrUpdate(path string, content []byte) (err error) {
	var file *os.File

	if is, _ := IsFolder(path); is {
		return fmt.Errorf("create file fails. The path is a folder")
	}

	if !Exists(path) {
		err = MkdirIfNotExists(filepath.Dir(path))
		if err != nil {
			return errors.Wrapf(err, "creating folders for file: %s", path)
		}

		file, err = os.Create(path)
		if err != nil {
			return errors.Wrapf(err, "creating new file %s", path)
		}
		log.WithField("path", path).Debugf("file created")
	} else {
		file, err = os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0755)
		if err != nil {
			return errors.Wrapf(err, "opening file %s", path)
		}
		log.WithField("path", path).Debugf("file opened")
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		return errors.Wrapf(err, "updating content to file %s", path)
	}

	_, err = file.Write(content)
	if err != nil {
		return errors.Wrapf(err, "writing content to file %s", path)
	}

	return nil
}

func Remove(path string) error {
	return os.Remove(path)
}

func GetContent(path string) ([]byte, error) {
	return nil, nil
}

func GetEmpty(path string) (*proto.File, error) {
	if !strings.HasPrefix(path, "./") {
		path = "./" + path
	}

	return proto.NewFile(path), nil
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

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func IsFolder(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, errors.Wrapf(err, "isFolder: getting stats for %s", path)
	}

	return fi.Mode().IsDir(), nil
}

func listFolders(path string) ([]string, error) {
	folders := make([]string, 0)
	elements, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, "listing folders from %s", path)
	}

	for _, element := range elements {
		if !element.IsDir() {
			continue
		}

		folders = append(folders, element.Name())
	}
	return folders, nil
}

func GetListFileRecursive(path string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if is, _ := IsFolder(path); !is {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrapf(err, "listing files recursively in %s", path)
	}

	return files, nil
}

func MkdirIfNotExists(path string) error {
	if !strings.HasPrefix(path, "/") {
		path = "./" + path
	}
	if Exists(path) {
		return nil
	}
	return os.MkdirAll(path, 0755)
}
