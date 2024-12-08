package main

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	env := Environment{}
	for _, e := range entries {
		fi, _ := e.Info()
		if fi.Size() == 0 {
			env[e.Name()] = EnvValue{"", true}
			continue
		}
		fp := filepath.Join(dir, e.Name())
		file, err := os.Open(fp)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			envVal := scanner.Text()
			if strings.Contains(envVal, "=") {
				return nil, errors.New("unsupported file: env contains =")
			}
			envVal = strings.Split(envVal, string('\000'))[0] // Обрезать строку после терминального нуля (если он есть).
			envVal = strings.Trim(envVal, " 	")
			env[e.Name()] = EnvValue{envVal, false}
			break // Прочитать только первую строку.
		}
		file.Close()
	}
	return env, nil
}
