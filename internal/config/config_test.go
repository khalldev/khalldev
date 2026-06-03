package config

import (
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	t.Setenv("GITHUB_USER", "")
	t.Setenv("MEDIUM_USER", "")
	c := Load()
	if c.GitHubUser != "khalldev" {
		t.Errorf("GitHubUser default = %q, want khalldev", c.GitHubUser)
	}
	if c.MediumUser != "khalcraft" {
		t.Errorf("MediumUser default = %q, want khalcraft", c.MediumUser)
	}
	if c.MaxWatched != 8 || c.MaxArticles != 10 {
		t.Errorf("limits wrong: watched=%d articles=%d", c.MaxWatched, c.MaxArticles)
	}
}

func TestLoadOverride(t *testing.T) {
	t.Setenv("GITHUB_USER", "alice")
	t.Setenv("MEDIUM_USER", "bob")
	c := Load()
	if c.GitHubUser != "alice" || c.MediumUser != "bob" {
		t.Errorf("override failed: %+v", c)
	}
}
