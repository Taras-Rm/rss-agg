package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Taras-Rm/rss-agg/internal/database"
	"github.com/google/uuid"
)

func startScriping(queries *database.Queries, cuncurrency int, interval time.Duration) {
	log.Printf("Start scriping cuncurrency %v, with interval %v", cuncurrency, interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

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
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Printf("Failed parse time %v with error %v", item.PubDate, err)
			continue
		}

		_, err = queries.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			Url:         item.Link,
			PublishedAt: pubAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			fmt.Printf("Failed create post %v with error %v", item.Title, err)
			return
		}
	}
	log.Printf("Feed %v collected, posts %v found", rssFeeds.Chanel.Title, len(rssFeeds.Chanel.Item))
}
