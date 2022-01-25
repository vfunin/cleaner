package remover

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/vfunin/cleaner/internal/finder"
)

// GetDuplicatesCount - find all files from []finder.File where finder.File.Delete is true
func GetDuplicatesCount(files []*finder.File) int {
	cnt := 0

	for _, file := range files {
		if !file.Delete {
			continue
		}
		cnt++
	}

	return cnt
}

// RemoveDuplicates - removes file from []finder.File if finder.File.Delete is true
func RemoveDuplicates(files []*finder.File) error {
	wg := sync.WaitGroup{}

	for _, file := range files {
		if !file.Delete {
			continue
		}

		errChan := make(chan error)

		wg.Add(1)

		go func() {
			defer wg.Done()
			remove(file.Path, errChan)
		}()

		err := <-errChan

		if err != nil {
			return err
		}

		wg.Wait()
	}

	return nil
}

func remove(file string, c chan<- error) {
	err := os.Remove(file)
	if err != nil {
		c <- err
	}

	log.Info("the file was deleted successfully: ", file)
	close(c)
}
