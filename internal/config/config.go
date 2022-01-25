package config

import (
	"flag"
	"fmt"
)

const (
	defaultDir         = "./files/"
	defaultHelp        = false
	defaultForceDelete = false
	defaultFake        = false
)

// Configuration - main config structure
type Configuration struct {
	Dir         string
	ForceDelete bool
	NeedHelp    bool
	NeedFake    bool
}

// Load - loads configuration from flags
func Load() (c Configuration, err error) {
	flag.StringVar(&c.Dir, "d", defaultDir, "directory")
	flag.BoolVar(&c.ForceDelete, "f", defaultForceDelete, "force delete duplicates")
	flag.BoolVar(&c.NeedFake, "g", defaultFake, "fill directory with fake files")
	flag.BoolVar(&c.NeedHelp, "h", defaultHelp, "print help")
	flag.Parse()

	return
}

// ShowHelp - prints app help info
func ShowHelp() {
	fmt.Println(`
About:
Cleaner is a program for removing duplicate files in a directory and subdirectories.

Examples of using:
$ ./bin/cleaner -g # will create a directory "files" in current directory with files and duplicates
$ ./bin/cleaner -d /files # will list all duplicates in directory "files"
$ ./bin/cleaner -f /files -f # will remove all duplicates in directory "files"`)
}
