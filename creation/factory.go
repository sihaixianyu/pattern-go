package creation

import (
	"fmt"
	"io"
	"os"
)

type Store interface {
	Open(string) error
	ReadToStdout() string
	WriteToStore(string)
	Close() error
	ResetFilePtr()
}

type StoreType int

const (
	DiskStoreType StoreType = 1 << iota
	TempStoreType
)

func NewStore(t StoreType) Store {
	switch t {
	case TempStoreType:
		return newTempStore()
	case DiskStoreType:
		return newDiskStore()
	default:
		return nil
	}
}

// TempStore implements Store
type TempStore struct {
	fileName string
	store    *os.File
}

func newTempStore() *TempStore {
	return &TempStore{}
}

func (s *TempStore) Open(name string) error {
	var err error
	s.fileName = name
	s.store, err = os.Create(name)

	return err
}

func (s *TempStore) Close() error {
	err := s.store.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("File Descriptor: \"%v\" has been Closed!\n", s.store)

	err = os.Remove(s.fileName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("File Nmae: \"%s\" has been Removed!\n", s.fileName)

	return err
}

func (s *TempStore) WriteToStore(content string) {
	_, err := s.store.Write([]byte(content))
	if err != nil {
		panic(err)
	}

	fmt.Printf("WriteToStore:\n\t\"%s\" to file: %s\n", content, s.fileName)
}

func (s *TempStore) ReadToStdout() string {
	res := make([]byte, 32)
	n, err := s.store.Read(res)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ReadToStdout:\n\t\"%s\" from file: %s\n", res[:n], s.fileName)

	return string(res)
}

func (s *TempStore) ResetFilePtr() {
	_, err := s.store.Seek(0, 0)
	if err != nil {
		panic(err)
	}
}

// * DiskStore is Unimplemented
type DiskStore struct {
	store io.ReadWriteCloser
}

func newDiskStore() *DiskStore {
	return &DiskStore{}
}

func (s *DiskStore) Open(name string) error {
	var err error
	s.store, err = os.Open(name)

	return err
}

func (s *DiskStore) Close() error {
	return s.store.Close()
}

func (s *DiskStore) ReadToStdout() string {
	return ""
}

func (s *DiskStore) WriteToStore(content string) {}

func (s *DiskStore) ResetFilePtr() {

}
