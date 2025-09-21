package feeds

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/frontendninja10/blog-aggregator/internal/app"
	"github.com/frontendninja10/blog-aggregator/internal/database"
	"github.com/frontendninja10/blog-aggregator/pkg/rss"
	"github.com/google/uuid"
)

func scrapeFeeds(s *app.State, user database.User) {
	feed, err := s.DB.GetNextFollowedFeedToFetch(context.Background(), user.ID)
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}

	_, err = s.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	fetchedFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Println(err)
	}

	for _, item := range fetchedFeed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time: t,
				Valid: true,
			}
		}
		postArgs := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid: true,
			},
			PublishedAt: publishedAt,
			FeedID: feed.ID,
			Url: item.Link,
		}
		_, err := s.DB.CreatePost(context.Background(), postArgs)

		if err != nil {
			log.Printf("couldn't create post: %v", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(fetchedFeed.Channel.Item))
	fmt.Println("==============================================================")
}

func AggregateFeeds(s *app.State, cmd app.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Printf("Collecting feeds every %v", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s, user)
	}
}