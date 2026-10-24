package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/FilledEther/rss_reader/internal/database"
	"github.com/google/uuid"
)

/*
	A scraper to scrap posts from rss feeds followed by the user.
	It is a function that is required to be working in the background for the entire duration of the server being live.
*/

func startScraping(db *database.Queries, timebetweenRequest time.Duration, concurrentFeeds int) {
	log.Printf("Scraping on %v goroutines every %s", concurrentFeeds, timebetweenRequest)
	ticker := time.NewTicker(timebetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedstoFetch(context.Background(), int32(concurrentFeeds))

		if err != nil {
			log.Println("Error in fetching:", err)
			continue //rather than returning because this worker should work in background
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeeds(db, wg, feed)

		}
		wg.Wait()
	}
}

func scrapeFeeds(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched", err)
		return
	}

	rssfeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error converting feed url to rss format.")
		return
	}

	for _, item := range rssfeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := parseDate(item.PubDate)

		if err != nil {
			log.Printf("couldn't parse date %v with err %v", item.PubDate, err)
		}

		_, err = db.CreatePost(context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: description,
				PublishedAt: pubAt,
				Url:         item.Link,
				FeedID:      feed.ID,
			})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("error fetching ", err)
		}
		log.Println("Found post", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssfeed.Channel.Item))
}

func parseDate(date string) (time.Time, error) {
	formats := []string{
		time.RFC1123Z,         // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,          // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC3339,          // "2006-01-02T15:04:05Z07:00"
		"2006-01-02 15:04:05", // Custom format: "YYYY-MM-DD HH:mm:ss"
	}
	var err error

	for _, format := range formats {
		parsedDate, err := time.Parse(format, date)
		if err == nil {
			return parsedDate, nil
		}
	}

	return time.Time{}, err
}
