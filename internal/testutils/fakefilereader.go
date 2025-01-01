package testutils

type FakeFileReader struct {
	data []byte
}

func NewFakeFileReader(data []byte) *FakeFileReader {
	return &FakeFileReader{data: data}
}

func (ffr *FakeFileReader) ReadData() ([]byte, error) {
	return ffr.data, nil
}
