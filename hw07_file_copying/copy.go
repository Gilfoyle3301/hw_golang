package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
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

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("file %s is not a regular file", fromPath)
	}

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}
	if _, err = file.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	writer, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer writer.Close()
	if limit == 0 {
		limit = fileInfo.Size()
		bar := progressbar.DefaultBytes(limit, "copy")
		_, err = io.Copy(io.MultiWriter(writer, bar), file)
	} else {
		bar := progressbar.DefaultBytes(limit, "copy")
		_, err = io.CopyN(io.MultiWriter(writer, bar), file, limit)
	}
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}
