
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

## Demo

![Demo](http://g.recordit.co/v5oS1juTje.gif)
# Usage

```
$ onlabs -h
Usage:
  onlabs  images    [--verbose|-v]
  onlabs  servers   [--verbose|-v]
  onlabs  volumes   [--verbose|-v]
  onlabs  snapshots [--verbose|-v]
  onlabs  ips       [--verbose|-v]
  onlabs  actions --server=IMAGEID  [--verbose|-v]
  onlabs  reboot --server=IMAGEID  [--verbose|-v]

Options:
  -h --help         this message
  -v --verbose      verbose mode
```

## Authentication

You need **token** to authorize the tool to call the API. You can manage the
tokens on the [credentials](https://cloud.online.net/#/credentials) page. Once
you have it set it as an environment variable:

```
export ONLINE_TOKEN=a0635283-123a-b456-90cd-0123456abcdef
```

## Bash tab completion

If you want **tab completition**, put this into your `.profile` / `.bash_profile`

```
complete -W "actions images servers volumes snapshots reboot ips --server --version --help" onlabs
```
