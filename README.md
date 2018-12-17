# birthtime-rename

[![Build Status](https://travis-ci.com/hioki-daichi/birthtime-rename.svg?branch=master)](https://travis-ci.com/hioki-daichi/birthtime-rename)
[![codecov](https://codecov.io/gh/hioki-daichi/birthtime-rename/branch/master/graph/badge.svg)](https://codecov.io/gh/hioki-daichi/birthtime-rename)
[![Go Report Card](https://goreportcard.com/badge/github.com/hioki-daichi/birthtime-rename)](https://goreportcard.com/report/github.com/hioki-daichi/birthtime-rename)

`birthtime-rename` is a command that can rename files to the birth time of the file.

## How to try

```shell-session
$ make build
go build -o birthtime-rename -v

$ tree testdata/
testdata/
└── a
    ├── a1.txt
    ├── a2.txt
    └── b
        ├── b1.txt
        └── b2.txt

$ ./birthtime-rename testdata/

$ tree testdata/
testdata/
└── a
    ├── 2018-12-17-10-06-42-001.txt
    ├── 2018-12-17-10-06-42-002.txt
    └── b
        ├── 2018-12-17-10-06-42-001.txt
        └── 2018-12-17-10-06-42-002.txt
```
