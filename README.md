
[![GoDoc](https://godoc.org/github.com/lalyos/onlabs?status.png)](https://godoc.org/github.com/lalyos/onlabs)
[![Travis](https://travis-ci.org/lalyos/onlabs.svg?branch=master)](https://travis-ci.org/lalyos/onlabs)

This project is a go library for [Online Labs](https://cloud.online.net/).
It also contains a command-line tool, which uses the [Online Labs API](https://doc.cloud.online.net/api/)


# Download

Linux/Darwin/Windows binaries are released at [github](https://github.com/lalyos/onlabs/releases/latest)

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
  onlabs  actions   --server=SERVERID  [--verbose|-v]
  onlabs  reboot    --server=SERVERID  [--verbose|-v]

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

If you want **tab completition**, use the `scripts/onlabs_completion`. You might
copy its content into your `.profile` / `.bash_profile`

Or if you want to use it in your actual shell:
```
curl -Ls j.mp/onlabs_completion | bash
```

## Install from source

If you want to hack around, you need go installed, then the usual
`go get` will install it to `$GOPATH/bin`

```
go get github.com/lalyos/onlabs/...
```
## Instal on OSX

The most popular package manager on mac is [brew](http://brew.sh).
Onlab is installable via:

```
brew cask install onlabs
```

If you are missing the [cask](http://caskroom.io) plugin, get it by:
```
brew install caskroom/cask/brew-cask
```

