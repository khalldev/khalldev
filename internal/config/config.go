package config

import "os"

type Config struct {
	GitHubUser  string
	MediumUser  string
	GitHubToken string
	MaxWatched  int
	MaxArticles int
}

func Load() Config {
	return Config{
		GitHubUser:  env("GITHUB_USER", "khalldev"),
		MediumUser:  env("MEDIUM_USER", "khalkhalkhal"),
		GitHubToken: os.Getenv("GITHUB_TOKEN"),
		MaxWatched:  8,
		MaxArticles: 10,
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
