package infrastructures

import (
	"io/ioutil"
	"os"

	"github.com/kou-pg-0131/circle-env/src/interfaces/gateways"
)

type FileSystem struct{}

func NewFileSystem() *FileSystem {
	return new(FileSystem)
}

func (fs *FileSystem) Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func (fs *FileSystem) IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (fs *FileSystem) Mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (fs *FileSystem) Create(path string) (gateways.IFile, error) {
	return os.Create(path)
}
