package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/muhamadhaspin/rssagg/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scarping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched:", err)
		return
	}

	rssFeed, err := FeedFromURL(feed.Url)
	if err != nil {
		log.Println("error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		desc := sql.NullString{
			String: item.Description,
			Valid:  true,
		}
		if item.Description != "" {
			desc.String = item.Description
			desc.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldn't parse publish date %v with error %v", item.PubDate, err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Description: desc,
			Url:         item.Link,
			FeedID:      feed.ID,
			PublishedAt: pubAt,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}

			log.Println("Couldn't create post:", err)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
