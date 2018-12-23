# birthtime-rename

[![Build Status](https://travis-ci.com/hioki-daichi/birthtime-rename.svg?branch=master)](https://travis-ci.com/hioki-daichi/birthtime-rename)
[![codecov](https://codecov.io/gh/hioki-daichi/birthtime-rename/branch/master/graph/badge.svg)](https://codecov.io/gh/hioki-daichi/birthtime-rename)
[![Go Report Card](https://goreportcard.com/badge/github.com/hioki-daichi/birthtime-rename)](https://goreportcard.com/report/github.com/hioki-daichi/birthtime-rename)

`birthtime-rename` is a command that can rename files to the birth time of the file.

## How to try

Try using `./testdata/` in the repository.

```
$ tree ./testdata/
./testdata/
└── a
    ├── a1.txt
    ├── a2.txt
    └── b
        ├── b1.txt
        └── b2.txt

2 directories, 4 files
```

Build with `make build`.

```
$ make build
go build -o birthtime-rename -v
```

Check with `--dry-run` option.

```
$ ./birthtime-rename --dry-run ./testdata/
testdata/a/a1.txt -> testdata/a/2018-12-23-16-20-12-e6d9715.txt
testdata/a/a2.txt -> testdata/a/2018-12-23-16-20-12-5447c6b.txt
testdata/a/b/b1.txt -> testdata/a/b/2018-12-23-16-20-12-b551771.txt
testdata/a/b/b2.txt -> testdata/a/b/2018-12-23-16-20-12-9d4b380.txt

$ tree ./testdata/
./testdata/
└── a
    ├── a1.txt
    ├── a2.txt
    └── b
        ├── b1.txt
        └── b2.txt

2 directories, 4 files
```

Actually rename.

```
$ ./birthtime-rename ./testdata/
testdata/a/a1.txt -> testdata/a/2018-12-23-16-20-12-e6d9715.txt
testdata/a/a2.txt -> testdata/a/2018-12-23-16-20-12-5447c6b.txt
testdata/a/b/b1.txt -> testdata/a/b/2018-12-23-16-20-12-b551771.txt
testdata/a/b/b2.txt -> testdata/a/b/2018-12-23-16-20-12-9d4b380.txt

$ tree ./testdata/
./testdata/
└── a
    ├── 2018-12-23-16-20-12-5447c6b.txt
    ├── 2018-12-23-16-20-12-e6d9715.txt
    └── b
        ├── 2018-12-23-16-20-12-9d4b380.txt
        └── 2018-12-23-16-20-12-b551771.txt

2 directories, 4 files
```
