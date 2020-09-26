package main

import (
	"flag"

	"github.com/moson-mo/mnserver/server"
)

func main() {
	url := flag.String("feed-url", "https://forum.manjaro.org/c/announcements.rss;https://forum.manjaro.org/c/arm/releases.rss;https://forum.manjaro.org/c/arm/stable-updates.rss;https://forum.manjaro.org/c/arm/testing-updates.rss;https://forum.manjaro.org/c/arm/unstable-updates.rss;https://forum.manjaro.org/c/arm/news.rss", "The RSS feed URL(s) separated by semicolon")
	interval := flag.Int("refresh-interval", 600, "The interval (in seconds) in which we check for new articles from the feed.")

	flag.Parse()

	server.Start(*url, *interval)
}
