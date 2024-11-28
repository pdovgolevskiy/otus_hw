package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"

	//nolint:depguard
	"github.com/stretchr/testify/require"
)

func equal(file1 string, file2 string) bool {
	f1, err1 := os.ReadFile(file1)

	if err1 != nil {
		log.Fatal(err1)
	}
	f2, err2 := os.ReadFile(file2)

	if err2 != nil {
		log.Fatal(err2)
	}
	return bytes.Equal(f1, f2)
}

func TestCopy(t *testing.T) {
	// перевод тестов из sh.
	const outputFilePath = "out.txt"
	defer os.Remove(outputFilePath)
	Copy("testdata/input.txt", outputFilePath, 0, 0)
	require.True(t, equal(outputFilePath, "testdata/out_offset0_limit0.txt"))
	os.Remove(outputFilePath)

	Copy("testdata/input.txt", outputFilePath, 0, 10)
	require.True(t, equal(outputFilePath, "testdata/out_offset0_limit10.txt"))
	os.Remove(outputFilePath)

	Copy("testdata/input.txt", outputFilePath, 0, 1000)
	require.True(t, equal(outputFilePath, "testdata/out_offset0_limit1000.txt"))
	os.Remove(outputFilePath)
	// Тесты на ошибки.
	err := Copy("file_not_exists.txt", outputFilePath, 0, 0)
	require.True(t, errors.Is(err, os.ErrNotExist))
	os.Remove(outputFilePath)

	err2 := Copy("testdata/input.txt", outputFilePath, 2<<31, 0)
	require.Equal(t, ErrOffsetExceedsFileSize, err2)
	os.Remove(outputFilePath)
}
