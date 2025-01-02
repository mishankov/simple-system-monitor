package fsutils

import "os"

type FileReader struct {
	path string
}

func NewFileReader(path string) *FileReader {
	return &FileReader{path: path}
}

func (fr *FileReader) ReadData() ([]byte, error) {
	return os.ReadFile(fr.path)
}
