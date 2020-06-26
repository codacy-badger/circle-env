package gateways

type IFile interface {
	Write(b []byte) (int, error)
	Close() error
}

type IFileSystem interface {
	Read(path string) ([]byte, error)
	IsExists(path string) bool
	Mkdir(path string) error
	Create(path string) (IFile, error)
}
