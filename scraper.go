package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Taras-Rm/rss-agg/internal/database"
)

func startScriping(queries *database.Queries, cuncurrency int, interval time.Duration) {
	log.Printf("Start scriping cuncurrency %v, with interval %v", cuncurrency, interval)

	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
		log.Println("Start scripe feed")
		feeds, err := queries.GetNextFeedsToFetch(context.Background(), int32(cuncurrency))
		if err != nil {
			fmt.Println("Error fetching feeds")
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(queries, wg, feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(queries *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	rssFeeds, err := urlToFeed(feed.Url)
	if err != nil {
		fmt.Printf("Error fetching feed: %v", err)
		return
	}

	_, err = queries.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Println("Error marking feed as fetched")
		return
	}

	for _, item := range rssFeeds.Chanel.Item {
		log.Printf("Found post: %v", item.Title)
	}
	log.Printf("Feed %v collected, posts %v found", rssFeeds.Chanel.Title, len(rssFeeds.Chanel.Item))
}
