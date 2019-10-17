package main

import (
	"flag"

	"github.com/moson-mo/mnserver/server"
)

func main() {
	url := flag.String("feed-url", "https://forum.manjaro.org/c/announcements.rss", "The RSS feed URL")
	interval := flag.Int("refresh-interval", 600, "The interval (in seconds) in which we check for new articles from the feed.")

	flag.Parse()

	server.Start(*url, *interval)
}
