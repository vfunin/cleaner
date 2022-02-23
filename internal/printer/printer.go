package printer

import (
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/vfunin/cleaner/internal/finder"
)

// ShowDuplicates - shows all duplicates
func ShowDuplicates(files []*finder.File) {
	wg := &sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)

		go func(file *finder.File) {
			defer wg.Done()
			showDuplicatesPerFile(file, files)
		}(file)
	}

	wg.Wait()
}

// showDuplicatesPerFile - finds and shows duplicates for files
func showDuplicatesPerFile(file *finder.File, files []*finder.File) {
	var duplicates []string

	for _, f := range files {
		if f.Name == file.Name && f.Path != file.Path && f.Size == file.Size {
			duplicates = append(duplicates, f.Path)
		}
	}

	if duplicates == nil {
		return
	}

	log.Info("found duplicates for file ", file.Path, ": ", strings.Join(duplicates, ","))
}
