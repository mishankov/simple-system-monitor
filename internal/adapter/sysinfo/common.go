package sysinfo

type DataReader interface {
	ReadData() ([]byte, error)
}
