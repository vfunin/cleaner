package main

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
	baseDir = "/files/"
	//duplicateFilesCnt = 1
	uniqueFilesCnt = 10
	nameLength     = 10
	contentLength  = 20
	maxDepth       = 5
	bytes          = 32
)

//type File struct {
//	Path    string
//	Content string
//}
//
//type DuplicateCounter struct {
//	sync.Mutex
//	cnt uint32
//}
//
//func (d *DuplicateCounter) Inc() {
//	d.Lock()
//	defer d.Unlock()
//	d.cnt++
//}
//
//func (d *DuplicateCounter) GetCnt() uint32 {
//	d.Lock()
//	defer d.Unlock()
//	return d.cnt
//}

func getRandomString(length uint8) (string, error) {
	randomBytes := make([]byte, bytes)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", errors.Wrap(err, "error in generation random string")
	}

	return strings.ToLower(base32.StdEncoding.EncodeToString(randomBytes)[:length]), nil
}

func main() {
	mRand.Seed(time.Now().UnixNano())

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir += baseDir
	err = createFilesAndDirectories(dir)
	log.Fatal(err)
}

func createFilesAndDirectories(baseDir string) (err error) {
	var (
		fName    string
		fContent string
		fPath    string
		depth    uint8
	)

	for i := 0; i < uniqueFilesCnt; i++ {
		depth = uint8(mRand.Intn(maxDepth)) //nolint:gosec
		fName, err = getRandomString(nameLength)

		if err != nil {
			return err
		}

		fName += ".txt"
		fContent, err = getRandomString(contentLength)

		if err != nil {
			return err
		}

		fPath, err = createDir(baseDir, depth)

		if err != nil {
			return err
		}

		err = CreateFile(fPath+fName, fContent)
	}

	return err
}

func createDir(path string, depth uint8) (string, error) {
	var i uint8
	for i = 0; i < depth; i++ {
		dir, err := getRandomString(maxDepth)
		if err != nil {
			return "", errors.Wrap(err, "error in dir creation")
		}

		path += dir + "/"
	}

	return path, os.MkdirAll(path, os.ModePerm)
}

func CreateFile(path string, content string) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	defer func() {
		err = errors.New("test")
	}()

	if _, err = f.WriteString(content); err != nil {
		return err
	}

	return f.Sync()
}
