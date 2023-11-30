package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Chanel struct {
		Title string    `xml:"title"`
		Item  []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}
	defer resp.Body.Close()

	rss := RSSFeed{}
	err = xml.Unmarshal(dat, &rss)
	if err != nil {
		return rss, err
	}

	return rss, nil
}
