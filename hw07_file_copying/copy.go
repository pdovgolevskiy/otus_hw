package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	const bufSize int = 1024
	file, err := os.Open(fromPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
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
	//TMP
	if limit == 0 || limit > sz {
		limit = sz
	}
	//readByteLen := int(sz - offset)
	bar := pb.New(int(limit)).SetUnits(pb.U_BYTES)
	bar.Start()
	newFile.Seek(offset, 0)
	writer := io.MultiWriter(newFile, bar)
	io.CopyN(writer, file, int64(limit))
	bar.Finish()
	return nil
}
