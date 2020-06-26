package infrastructures

import (
	"io/ioutil"
	"os"

	"github.com/kou-pg-0131/circle-env/src/interfaces/gateways"
)

// FileSystem ...
type FileSystem struct{}

// NewFileSystem ...
func NewFileSystem() *FileSystem {
	return new(FileSystem)
}

// ReadFile ...
func (fs *FileSystem) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// IsExists ...
func (fs *FileSystem) IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Mkdir ...
func (fs *FileSystem) Mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Create ...
func (fs *FileSystem) Create(path string) (gateways.IFile, error) {
	return os.Create(path)
}
