package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

var (
	errArgumentRequired              = errors.New("argument required")
	errOnlyOneArgumentCanBeSpecified = errors.New("only one argument can be specified")
)

func main() {
	if err := execute(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func execute(args []string) error {
	l := len(args)

	if l == 0 {
		return errArgumentRequired
	} else if l > 1 {
		return errOnlyOneArgumentCanBeSpecified
	}

	return filepath.Walk(args[0], walkFn)
}

func walkFn(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if fi.IsDir() {
		return nil
	}

	return rename(path, fi)
}

func rename(path string, fi os.FileInfo) error {
	newpath := genNewpath(path, fi)

	_, err := os.OpenFile(newpath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}

	err = os.Rename(path, newpath)
	if err != nil {
		return err
	}

	return nil
}

func genNewpath(path string, fi os.FileInfo) string {
	birthTime := getBirthTime(fi)

	fmtBtime := birthTime.Format("2006-01-02-15-04-05")
	ext := filepath.Ext(path)

	return filepath.Join(filepath.Dir(path), fmtBtime+ext)
}
