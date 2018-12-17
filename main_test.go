package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func Test_execute_noSuchFileOrDirectory(t *testing.T) {
	expected := "lstat non-existent-path: no such file or directory"

	err := execute([]string{"non-existent-path"})

	if err == nil {
		t.Fatal("unexpectedly err is nil")
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf(`unexpected err: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func Test_execute(t *testing.T) {
	path, clean := copyTestdataToTempDir(t)
	defer clean()

	err := execute([]string{path})
	if err != nil {
		t.Fatalf("err %s", err)
	}
}

func copyTestdataToTempDir(t *testing.T) (string, func()) {
	t.Helper()

	destdir, err := ioutil.TempDir("", "birthtime-rename")
	if err != nil {
		t.Fatalf("err %s", err)
	}

	err = filepath.Walk("testdata", func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destpath := filepath.Join(destdir, path)

		if fi.IsDir() {
			return os.Mkdir(destpath, 0755)
		}

		sourcefile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer sourcefile.Close()

		destfile, err := os.Create(destpath)
		if err != nil {
			return err
		}
		defer destfile.Close()

		_, err = io.Copy(destfile, sourcefile)
		return err
	})

	if err != nil {
		t.Fatalf("err %s", err)
	}

	return destdir, func() { os.Remove(destdir) }
}
