package faker

import (
	"crypto/rand"
	"encoding/base32"
	mRand "math/rand"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	filesDir          = "/files/"
	duplicateFilesCnt = 3
	filesCnt          = 10
	nameLength        = 10
	contentLength     = 20
	maxDepth          = 5
	bytes             = 32
)

type File struct {
	Path    string
	Content string
	Name    string
}

func getRandomString(length uint8) (string, error) {
	randomBytes := make([]byte, bytes)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", errors.Wrap(err, "error in generation random string")
	}

	return strings.ToLower(base32.StdEncoding.EncodeToString(randomBytes)[:length]), nil
}

func Run() {
	log.SetFormatter(&log.JSONFormatter{}) //nolint:exhaustivestruct
	log.Info("starting faker")
	mRand.Seed(time.Now().UnixNano())

	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	baseDir += filesDir
	if err = createFilesAndDirectories(baseDir); err != nil {
		log.Fatal(err)
	}

	log.Printf("faker done. %d files created", filesCnt+duplicateFilesCnt)
}

func createFilesAndDirectories(baseDir string) (err error) {
	maxDuplicates := duplicateFilesCnt

	var file *File

	for i := 0; i < filesCnt; i++ {
		if file, err = createDirectoryAndFile(baseDir, nil); err != nil {
			return
		}

		log.Printf("new file created: \"%s\"", file.Path)

		if maxDuplicates > 0 {
			if file, err = createDirectoryAndFile(baseDir, file); err != nil {
				return
			}

			log.Printf("new file duplicate created \"%s\"", file.Path)
			maxDuplicates--
		}
	}

	return err
}

func createDirectoryAndFile(baseDir string, existingFile *File) (*File, error) {
	var (
		file     File
		fName    string
		fContent string
		fPath    string
		depth    uint8
		err      error
	)

	depth = uint8(mRand.Intn(maxDepth)) //nolint:gosec

	if existingFile == nil {
		fName, err = getRandomString(nameLength)
		if err != nil {
			return nil, err
		}

		fName += ".txt"
	} else {
		fName = existingFile.Name
	}

	if existingFile == nil {
		fContent, err = getRandomString(contentLength)
	} else {
		fContent = existingFile.Content
	}

	if err != nil {
		return nil, err
	}

	fPath, err = createDir(baseDir, depth)

	if err != nil {
		return nil, err
	}

	file, err = CreateFile(fPath, fName, fContent)

	return &file, err
}

func createDir(path string, depth uint8) (string, error) {
	var i uint8
	for i = 0; i < depth; i++ {
		dir, err := getRandomString(maxDepth)
		if err != nil {
			return "", errors.Wrap(err, "error in generation dir name")
		}

		path += dir + "/"
	}

	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return "", errors.Wrap(err, "error in dir creation")
	}

	log.Printf("directory \"%s\" created", path)

	return path, err
}

func CreateFile(path string, name string, content string) (file File, err error) {
	f, err := os.Create(path + name)

	if err != nil {
		return
	}

	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		return
	}

	err = f.Sync()

	if err != nil {
		return
	}

	file.Name = name
	file.Path = path + name
	file.Content = content

	return
}
