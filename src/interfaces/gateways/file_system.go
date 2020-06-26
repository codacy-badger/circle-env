package gateways

// IFile ...
type IFile interface {
	Write(b []byte) (int, error)
	Close() error
}

// IFileSystem ...
type IFileSystem interface {
	ReadFile(path string) ([]byte, error)
	IsExists(path string) bool
	Mkdir(path string) error
	Create(path string) (IFile, error)
}
