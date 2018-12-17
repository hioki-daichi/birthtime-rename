package main

import (
	"os"
	"syscall"
	"time"
)

func getBirthTime(fi os.FileInfo) time.Time {
	return time.Unix(fi.Sys().(*syscall.Stat_t).Birthtimespec.Sec, 0)
}
