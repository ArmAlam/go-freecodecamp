package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/armalam/go-freecodecamp/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))

		if err != nil {
			fmt.Println("Error fetching feed ", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}

		wg.Wait()
	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {

	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Println("Error fetching feed ", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Println("Error fetching feed ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Fetched post", item.Title)
		// TODO: SAVE IN DB
	}

	log.Printf("Feed  %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}
