package main

import (
	"errors"
	"io"
	"os"

	//nolint:depguard
	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	sz := fi.Size()
	if offset > sz {
		return ErrUnsupportedFile
	}
	mode := fi.Mode()
	newFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer newFile.Close()
	os.Chmod(toPath, mode)
	if limit == 0 || limit > sz {
		limit = sz
	}
	bar := pb.New(int(limit)).SetUnits(pb.U_BYTES)
	bar.Start()
	file.Seek(offset, 0)
	writer := io.MultiWriter(newFile, bar)
	io.CopyN(writer, file, limit)
	bar.Finish()
	return nil
}
