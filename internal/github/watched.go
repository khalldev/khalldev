package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Watched struct {
	Name string
	URL  string
	Date string
}

type event struct {
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	Repo      struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
}

func FetchWatched(baseURL, user, token string, max int) ([]Watched, error) {
	url := fmt.Sprintf("%s/users/%s/events/public?per_page=100", baseURL, user)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch events: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("github status %d: %s", resp.StatusCode, string(body))
	}

	var events []event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("decode events: %w", err)
	}

	seen := make(map[string]bool)
	out := make([]Watched, 0, max)
	for _, e := range events {
		if e.Type != "WatchEvent" || seen[e.Repo.Name] {
			continue
		}
		seen[e.Repo.Name] = true
		out = append(out, Watched{
			Name: e.Repo.Name,
			URL:  "https://github.com/" + e.Repo.Name,
			Date: e.CreatedAt.UTC().Format("2006-01-02"),
		})
		if len(out) >= max {
			break
		}
	}
	return out, nil
}
