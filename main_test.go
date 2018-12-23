package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var actualErr error

func Test_main_ErrArgumentRequired(t *testing.T) {
	expected := errArgumentRequired

	logFatalFunc = func(v ...interface{}) {
		actualErr = v[0].(error)
	}

	os.Args = []string{"birthtime-rename"}

	main()

	if actualErr != expected {
		t.Errorf(`unexpected err: expected: "%s" actual: "%s"`, expected, actualErr)
	}
}

func Test_execute_ErrArgumentRequired(t *testing.T) {
	expected := errArgumentRequired.Error()

	err := execute([]string{"birthtime-rename"})
	if err == nil {
		t.Fatal("unexpectedly err is nil")
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf(`unexpected err: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func Test_execute_ErrOnlyOneArgumentCanBeSpecified(t *testing.T) {
	expected := errOnlyOneArgumentCanBeSpecified.Error()

	err := execute([]string{"birthtime-rename", "a", "b"})
	if err == nil {
		t.Fatal("unexpectedly err is nil")
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf(`unexpected err: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func Test_execute_NoSuchFileOrDirectory(t *testing.T) {
	expected := "lstat non-existent-path: no such file or directory"

	err := execute([]string{"birthtime-rename", "non-existent-path"})

	if err == nil {
		t.Fatal("unexpectedly err is nil")
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf(`unexpected err: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func Test_execute(t *testing.T) {
	getBirthTimeFunc = func(fi os.FileInfo) time.Time {
		loc, err := time.LoadLocation("UTC")
		if err != nil {
			panic(err)
		}

		return time.Unix(0, 0).In(loc)
	}

	dirname, clean := copyTestdataToTempDir(t)
	defer clean()

	paths := []string{
		filepath.Join(dirname, "testdata/a/1970-01-01-00-00-00-5447c6b.txt"),
		filepath.Join(dirname, "testdata/a/1970-01-01-00-00-00-e6d9715.txt"),
		filepath.Join(dirname, "testdata/a/b/1970-01-01-00-00-00-9d4b380.txt"),
		filepath.Join(dirname, "testdata/a/b/1970-01-01-00-00-00-b551771.txt"),
	}

	err := execute([]string{"birthtime-rename", dirname})
	if err != nil {
		t.Fatalf("err %s", err)
	}

	for _, path := range paths {
		if _, err = os.Stat(path); os.IsNotExist(err) {
			t.Fatalf("Unexpectedly a file or directory did not exist at the path: %s", path)
		}
	}

	// The result is the same even if it is executed twice.
	err = execute([]string{"birthtime-rename", dirname})
	if err != nil {
		t.Fatalf("err %s", err)
	}

	for _, path := range paths {
		if _, err = os.Stat(path); os.IsNotExist(err) {
			t.Fatalf("Unexpectedly a file or directory did not exist at the path: %s", path)
		}
	}
}

func Test_execute_Unreadable(t *testing.T) {
	dirname, clean := copyTestdataToTempDir(t)
	defer clean()

	path := filepath.Join(dirname, "unreadable")

	expected := "open " + path + ": permission denied"

	_, err := os.OpenFile(path, os.O_CREATE, 0)
	if err != nil {
		t.Fatalf("err %s", err)
	}

	err = execute([]string{"birthtime-rename", dirname})
	if err == nil {
		t.Fatal("unexpectedly err is nil")
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf(`unexpected err: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func Test_execute_RenameFailure(t *testing.T) {
	expectedErr := errors.New("failed to rename")

	renameFunc = func(oldpath, newpath string) error {
		return expectedErr
	}

	dirname, clean := copyTestdataToTempDir(t)
	defer clean()

	err := execute([]string{"birthtime-rename", dirname})
	if err != expectedErr {
		t.Errorf(`unexpected err: expected: "%s" actual: "%s"`, expectedErr, err)
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
