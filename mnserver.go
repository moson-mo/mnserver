package main

import (
	"flag"

	"github.com/moson-mo/mnserver/server"
)

func main() {
	url := flag.String("feed-url", "https://forum.manjaro.org/c/announcements.rss;https://forum.manjaro.org/c/manjaro-arm/announcements.rss;https://forum.manjaro.org/c/manjaro-arm/releases.rss;https://forum.manjaro.org/c/manjaro-arm/arm-stable-updates.rss;https://forum.manjaro.org/c/manjaro-arm/arm-testing-updates.rss", "The RSS feed URL")
	twitterURL := flag.String("twitter-url", "https://twitter.com/manjarolinux", "The manjarolinux twitter URL")
	interval := flag.Int("refresh-interval", 600, "The interval (in seconds) in which we check for new articles from the feed.")

	flag.Parse()

	server.Start(*url, *twitterURL, *interval)
}
