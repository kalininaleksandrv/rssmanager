package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"golang.org/x/text/encoding/charmap"
)

type RssFeed struct {
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Items       []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

func urlToFeed(url string, wg *sync.WaitGroup) (RssFeed, error) {
	defer wg.Done()
	start := time.Now()
	fmt.Printf("---> Start fetching URL %s at %v\n", url, start)
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	rssFeed := RssFeed{}
	resp, err := httpClient.Get(url)
	if (err != nil) {
		return rssFeed, err
	}
	defer resp.Body.Close()
	payload, err := io.ReadAll(resp.Body)
	if (err != nil) {
		return rssFeed, err
	}

	reader := bytes.NewReader(payload)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "ISO-8859-1":
			return charmap.ISO8859_1.NewDecoder().Reader(input), nil
		case "UTF-8":
			return input, nil
		default:
			return nil, fmt.Errorf("unsupported charset: %s", charset)
		}
	}
	err = decoder.Decode(&rssFeed)
	if (err != nil) {
		return rssFeed, err
	}
	fmt.Printf("Fetched %s in %v\n", url, time.Since(start))
	return rssFeed, nil
}