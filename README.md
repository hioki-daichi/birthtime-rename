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
    ├── 2018-12-23-14-34-14-5447c6b.txt
    ├── 2018-12-23-14-34-14-e6d9715.txt
    └── b
        ├── 2018-12-23-14-34-14-9d4b380.txt
        └── 2018-12-23-14-34-14-b551771.txt
```
