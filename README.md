# mnserver - REST API service providing Manjaro Linux announcements

A small http server which delivers Manjaro announcements for the systray application (see matray repo).
It parses the Manjaro Linux announcements RSS feed in a regular interval. Clients can grab news articles from the server via http post request.

## Installation

Make sure go is installed on your system.

Install with:
```
go get github.com/moson-mo/mnserver
```

Run with:
```
$(go env GOPATH)/bin/mnserver
```

For a system-wide installation, just copy the binary to e.g. /usr/local/bin
(or run install.sh)

## How to build

Run `go build` in the mnserver directory or use the build script `build.sh`

## Libraries used

* [gofeed](https://github.com/mmcdole/gofeed) - RSS/Atom feed parser
* [twitter-scraper](https://github.com/n0madic/twitter-scraper) - Twitter scraper

