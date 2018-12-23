package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var dryRun bool

var (
	errArgumentRequired              = errors.New("argument required")
	errOnlyOneArgumentCanBeSpecified = errors.New("only one argument can be specified")
)

func main() {
	if err := execute(os.Args); err != nil {
		log.Fatal(err)
	}
}

func execute(args []string) error {
	flg := flag.NewFlagSet(args[0], flag.ExitOnError)

	flg.BoolVar(&dryRun, "dry-run", false, "Print the execution result without executing it")

	flg.Parse(args[1:])

	ss := flg.Args()

	l := len(ss)
	if l == 0 {
		return errArgumentRequired
	} else if l > 1 {
		return errOnlyOneArgumentCanBeSpecified
	}

	return filepath.Walk(ss[0], walkFn)
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
	newpath, err := genNewpath(path, fi)
	if err != nil {
		return err
	}

	if !dryRun {
		_, err = os.OpenFile(newpath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
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
	}

	fmt.Println(path + " -> " + newpath)

	return nil
}

func genNewpath(path string, fi os.FileInfo) (string, error) {
	birthTime := getBirthTime(fi)

	fmtBtime := birthTime.Format("2006-01-02-15-04-05")

	suffix, err := genSuffix(path)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(path)

	return filepath.Join(filepath.Dir(path), fmtBtime+"-"+suffix+ext), nil
}

func genSuffix(path string) (string, error) {
	h := sha1.New()

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}

	b := h.Sum(nil)

	return hex.EncodeToString(b)[:7], nil
}
