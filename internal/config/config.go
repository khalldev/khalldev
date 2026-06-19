package config

import "os"

type Config struct {
	GitHubUser  string
	GitHubToken string
	MaxWatched  int
}

func Load() Config {
	return Config{
		GitHubUser:  env("GITHUB_USER", "khalldev"),
		GitHubToken: os.Getenv("GITHUB_TOKEN"),
		MaxWatched:  8,
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
