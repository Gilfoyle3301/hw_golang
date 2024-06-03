package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
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
	hashMap := make(Environment, 0)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() || strings.Contains(file.Name(), "=") {
			continue
		}
		openFile, err := os.Open(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		defer openFile.Close()
		reader := bufio.NewReader(openFile)
		readEnv, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				hashMap[file.Name()] = EnvValue{NeedRemove: true}
				continue
			}
			return nil, err
		}
		changeEnvLineTab := bytes.TrimRight(readEnv, " \t")
		changeEnvLineZero := bytes.ReplaceAll(changeEnvLineTab, []byte("\x00"), []byte("\n"))
		hashMap[file.Name()] = EnvValue{string(changeEnvLineZero), false}
	}
	return hashMap, nil
}
