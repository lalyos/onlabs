
[![GoDoc](https://godoc.org/github.com/lalyos/onlabs?status.png)](https://godoc.org/github.com/lalyos/onlabs)
[![Travis](https://travis-ci.org/lalyos/onlabs.svg?branch=master)](https://travis-ci.org/lalyos/onlabs)

This project is a go library for [Online Labs](https://cloud.online.net/).
It also contains a command-line tool, which uses the [Online Labs API](https://doc.cloud.online.net/api/)


# Download

Linux/Darwin/Windows binaries are released at [github](https://github.com/lalyos/onlabs/releases/latest)

## Install from source

If you have go installed, then the `go get` will install it to
`$GOPATH/bin`

```
go get github.com/lalyos/onlin/...
```

# Usage

```
$ onlabs -h

Usage:
        online  images [--verbose|-v]

Options:
  -h --help         this message
  -v --verbose      verbose mode
```
