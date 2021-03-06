// +build !linux

package file

import (
	"os"
)

func (m *memFile) Open() (*os.File, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	go func() {
		defer w.Close()
		w.Write(m.content)
	}()
	return r, nil
}
