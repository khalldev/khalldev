package main

import (
	"log"
	"os"

	"github.com/khalldev/readme-profile/internal/config"
	"github.com/khalldev/readme-profile/internal/github"
	"github.com/khalldev/readme-profile/internal/medium"
	"github.com/khalldev/readme-profile/internal/render"
)

const githubAPI = "https://api.github.com"

func main() {
	cfg := config.Load()

	watched, err := github.FetchWatched(githubAPI, cfg.GitHubUser, cfg.GitHubToken, cfg.MaxWatched)
	if err != nil {
		log.Printf("warn: github fetch failed: %v", err)
	}

	articles, err := medium.FetchArticles(medium.FeedURL(cfg.MediumUser), cfg.MaxArticles)
	if err != nil {
		log.Printf("warn: medium fetch failed: %v", err)
	}

	out, err := render.Render(render.Data{
		User:     cfg.GitHubUser,
		Watched:  watched,
		Articles: articles,
	})
	if err != nil {
		log.Fatalf("render: %v", err)
	}

	if err := os.WriteFile("README.md", []byte(out), 0o644); err != nil {
		log.Fatalf("write README.md: %v", err)
	}
	log.Printf("wrote README.md (%d watched, %d articles)", len(watched), len(articles))
}
