package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title		string `xml:"title"`
		Link		string `xml:"link"`
		Description	string `xml:"description"`
		Item		[]RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title		string `xml:"title"`
	Link		string `xml:"link"`
	Description string `xml:"description"`
	PubDate		string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
        return nil, fmt.Errorf("error reading response: %w", err)
    }

	var rssFeedRes RSSFeed
	err = xml.Unmarshal(data, &rssFeedRes)

	rssFeedRes.Channel.Title = html.UnescapeString(rssFeedRes.Channel.Title)
	rssFeedRes.Channel.Description = html.UnescapeString(rssFeedRes.Channel.Description)

	for i, item := range rssFeedRes.Channel.Item {
		rssFeedRes.Channel.Item[i].Title = html.UnescapeString(item.Title)
		rssFeedRes.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	if err != nil {
        return nil, fmt.Errorf("error unmarshalling response: %w", err)
    }


	return &rssFeedRes, nil
}