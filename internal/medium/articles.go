package medium

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Article struct {
	Title string
	Link  string
}

type rss struct {
	Channel struct {
		Items []struct {
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"item"`
	} `xml:"channel"`
}

func FetchArticles(feedURL string, max int) ([]Article, error) {
	resp, err := http.Get(feedURL)
	if err != nil {
		return nil, fmt.Errorf("fetch feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("medium status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	var feed rss
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, fmt.Errorf("parse rss: %w", err)
	}

	out := make([]Article, 0, max)
	for _, it := range feed.Channel.Items {
		if len(out) >= max {
			break
		}
		out = append(out, Article{Title: it.Title, Link: it.Link})
	}
	return out, nil
}

func FeedURL(user string) string {
	return "https://medium.com/feed/@" + user
}
