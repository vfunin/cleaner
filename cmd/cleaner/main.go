package main

import (
	"os"

	"github.com/vfunin/cleaner/internal/remover"

	log "github.com/sirupsen/logrus"
	"github.com/vfunin/cleaner/internal/config"
	"github.com/vfunin/cleaner/internal/faker"
	"github.com/vfunin/cleaner/internal/finder"
	"github.com/vfunin/cleaner/internal/printer"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{}) //nolint:exhaustivestruct

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if cfg.NeedHelp {
		config.ShowHelp()
		os.Exit(0)
	}

	if cfg.NeedFake {
		faker.Run()
		os.Exit(0)
	}

	if cfg.Dir == "" {
		log.Fatal("directory parameter shouldn't be empty")
	}

	filesCh, errCh := finder.GetAllFiles(cfg.Dir)
	select {
	case files := <-filesCh:
		if cfg.ForceDelete {
			log.Printf("found %d files for remove", remover.GetDuplicatesCount(files))
			err = remover.RemoveDuplicates(files)

			if err != nil {
				log.Fatal(err)
			}

			return
		}

		if remover.GetDuplicatesCount(files) == 0 {
			log.Info("found 0 files for remove")
		}

		printer.ShowDuplicates(files)
	case err = <-errCh:
		log.Fatal(err)
	}
}
