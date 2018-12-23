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

// for testing
var (
	getBirthTimeFunc = getBirthTime
	renameFunc       = os.Rename
	logFatalFunc     = log.Fatal
)

func main() {
	if err := execute(os.Args); err != nil {
		logFatalFunc(err)
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
	t := getBirthTimeFunc(fi)

	hexStr, err := genHexStrFromFile(path)
	if err != nil {
		return err
	}

	newpath := filepath.Join(
		filepath.Dir(path),
		t.Format("2006-01-02-15-04-05")+"-"+hexStr[:7]+filepath.Ext(path),
	)

	if !dryRun {
		err = renameFunc(path, newpath)
		if err != nil {
			return err
		}
	}

	fmt.Println(path + " -> " + newpath)

	return nil
}

func genHexStrFromFile(path string) (string, error) {
	h := sha1.New()

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, _ = io.Copy(h, f)

	return hex.EncodeToString(h.Sum(nil)), nil
}
