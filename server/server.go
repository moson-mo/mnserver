package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
)

type clientSettings struct {
	MaxItems   int      `json:"MaxItems"`
	Categories []string `json:"Categories"`
}

type article struct {
	GUID          string    `json:"GUID"`
	URL           string    `json:"URL"`
	Title         string    `json:"Title"`
	Category      string    `json:"Category"`
	PublishedDate time.Time `json:"PublishedDate"`
}

var (
	parser   *gofeed.Parser
	articles []article
	mut      sync.Mutex
)

// Start starts the service
func Start(feedURL string, refreshInterval int) {
	parser = gofeed.NewParser()
	go getNewsLoop(feedURL, time.Duration(refreshInterval)*time.Second)

	http.HandleFunc("/news", handleRequest)
	err := http.ListenAndServe(":10111", nil)
	if err != nil {
		print("Error starting http server: " + err.Error())
		return
	}
}

// handles the client request
func handleRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}

	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	cs := clientSettings{}
	err := d.Decode(&cs)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error parsing json: " + err.Error()))
		return
	}

	a := getArticlesWithCategories(cs.Categories)
	max := len(a)
	if cs.MaxItems < max {
		max = cs.MaxItems
	}

	ts := a[0:max]
	sort.Slice(ts, func(i, j int) bool {
		return ts[i].PublishedDate.Unix() < ts[j].PublishedDate.Unix()
	})
	b, err := json.MarshalIndent(&ts, "", "    ")
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error decoding json: " + err.Error()))
		return
	}

	print("Client fetched news")
	w.Write(b)
}

// returns a filtered list of articles and sorts by date
func getArticlesWithCategories(cats []string) []article {
	ret := []article{}
	for _, a := range articles {
		if containsCategory(a, cats) {
			ret = append(ret, a)
		}
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].PublishedDate.Unix() > ret[j].PublishedDate.Unix()
	})
	return ret
}

// checks if an article has certain categories
func containsCategory(a article, cats []string) bool {
	if len(cats) == 1 && cats[0] == "All" {
		return true
	}
	for _, cat := range cats {
		if a.Category == cat {
			return true
		}
	}
	return false
}

// runs news download function in loop after a certain interval
func getNewsLoop(url string, refreshInterval time.Duration) {
	getNews(url)

	for {
		<-time.After(refreshInterval)
		getNews(url)
	}
}

// downloads rss feed, adds articles to list and sends it to the connected clients
func getNews(url string) {
	feed, err := parser.ParseURL(url)
	if err != nil {
		print("Error parsing news: " + err.Error())
	}

	print("Getting news from feed...")

	for _, e := range feed.Items {
		if !articleExits(e.GUID) {
			a := article{
				GUID:          e.GUID,
				URL:           e.Link,
				Title:         e.Title,
				Category:      e.Categories[0],
				PublishedDate: *e.PublishedParsed,
			}

			mut.Lock()
			articles = append(articles, a)
			mut.Unlock()
			print("New article added: " + e.Title)
		}
	}
}

// checks if our list already contains an article
func articleExits(guid string) bool {
	for _, a := range articles {
		if a.GUID == guid {
			return true
		}
	}
	return false
}

// prints out a message with the current timestamp
func print(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 - 15:04:05:\t") + msg)
}
