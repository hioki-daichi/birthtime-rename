package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	errArgumentRequired              = errors.New("argument required")
	errOnlyOneArgumentCanBeSpecified = errors.New("only one argument can be specified")
)

func main() {
	err := execute(os.Args[1:])
	if err != nil {
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
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	count := 1

	for {
		newFp, err := os.OpenFile(genNewpath(path, fi, count), os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
		if err != nil {
			if os.IsExist(err) {
				count++
				continue
			}
			return err
		}

		_, err = io.Copy(newFp, fp)
		if err != nil {
			return err
		}

		err = os.Remove(path)
		if err != nil {
			return err
		}

		break
	}

	return nil
}

func genNewpath(path string, fi os.FileInfo, count int) string {
	birthTime := getBirthTime(fi)

	fmtBtime := birthTime.Format("2006-01-02-15-04-05")
	fmtCount := fmt.Sprintf("-%03d", count)
	ext := filepath.Ext(path)

	return filepath.Join(filepath.Dir(path), fmtBtime+fmtCount+ext)
}
