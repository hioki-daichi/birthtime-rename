package main

import (
	"os"
	"testing"
)

func Test_getBirthTime(t *testing.T) {
	fi, err := os.Stat("./testdata/a/a1.txt")
	if err != nil {
		t.Fatalf("err %s", err)
	}

	_ = getBirthTime(fi)
}
