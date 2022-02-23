package finder

import (
	"fmt"
	"os"
	"path/filepath"
)

// File - main file structure
type File struct {
	Name   string
	Path   string
	Size   int64
	Delete bool
}

// GetAllFiles - returns two channels: files in dir and error
func GetAllFiles(dir string) (chan []*File, chan error) {
	var files []*File

	m := make(map[string]string)

	filesCh := make(chan []*File)
	errCh := make(chan error)

	go func() {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, _ error) error {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				errCh <- err

				return nil
			}

			if info.IsDir() {
				return nil
			}

			key := fmt.Sprintf("%v:%d", info.Name(), info.Size())

			f := File{
				Name:   info.Name(),
				Path:   path,
				Size:   info.Size(),
				Delete: true,
			}

			_, ok := m[key]
			if !ok {
				m[key] = path
				f.Delete = false
			}

			files = append(files, &f)

			return nil
		})
		if err != nil {
			errCh <- err

			return
		}
		filesCh <- files

		defer close(filesCh)
		defer close(errCh)
	}()

	return filesCh, errCh
}
