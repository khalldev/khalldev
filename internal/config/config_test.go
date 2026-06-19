package config

import (
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	t.Setenv("GITHUB_USER", "")
	c := Load()
	if c.GitHubUser != "khalldev" {
		t.Errorf("GitHubUser default = %q, want khalldev", c.GitHubUser)
	}
	if c.MaxWatched != 8 {
		t.Errorf("limits wrong: watched=%d", c.MaxWatched)
	}
}

func TestLoadOverride(t *testing.T) {
	t.Setenv("GITHUB_USER", "alice")
	c := Load()
	if c.GitHubUser != "alice" {
		t.Errorf("override failed: %+v", c)
	}
}
