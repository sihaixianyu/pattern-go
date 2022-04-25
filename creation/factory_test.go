package creation

import (
	"testing"
)

func TestNewStore(t *testing.T) {
	s := NewStore(TempStoreType)

	err := s.Open("temp")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	s.WriteToStore("Hello World!")
	s.ResetFilePtr()
	s.ReadToStdout()
}
