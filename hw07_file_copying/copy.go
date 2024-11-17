package main

import (
	"errors"
	"io"
	"log"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	const bufSize int = 1024
	file, err := os.Open(fromPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	sz := fi.Size()
	if offset > sz {
		log.Fatal(ErrUnsupportedFile)
	}
	//
	mode := fi.Mode()
	newFile, err := os.Create(toPath)
	if err != nil {
		log.Fatal(err)
	}
	os.Chmod(toPath, mode)
	//
	//TMP
	if limit == 0 || limit > sz {
		limit = sz
	}

	for offset < sz {
		io.CopyN(newFile, file, int64(bufSize))
		// read, err := file.Read(buf[offset:])
		// offset += int64(read)
		// if err == io.EOF {
		// // что если не дочитали ?
		// break
		// }
		// if err != nil {
		// log.Panicf("failed to read: %v", err)
		// }
	}
	// Place your code here.
	return nil
}
