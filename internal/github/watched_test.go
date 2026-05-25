package github

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFetchWatched(t *testing.T) {
	fixture, err := os.ReadFile("testdata/events.json")
	if err != nil {
		t.Fatal(err)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/khalldev/events/public" {
			t.Errorf("unexpected path %q", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer tok" {
			t.Errorf("missing auth header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(fixture)
	}))
	defer srv.Close()

	got, err := FetchWatched(srv.URL, "khalldev", "tok", 8)
	if err != nil {
		t.Fatal(err)
	}
	want := []Watched{
		{Name: "BenEmdon/CenteredCollectionView", URL: "https://github.com/BenEmdon/CenteredCollectionView", Date: "2026-05-25"},
		{Name: "kageroumado/phosphene", URL: "https://github.com/kageroumado/phosphene", Date: "2026-05-22"},
		{Name: "pystardust/ani-cli", URL: "https://github.com/pystardust/ani-cli", Date: "2026-05-11"},
	}
	if len(got) != len(want) {
		t.Fatalf("got %d items, want %d: %+v", len(got), len(want), got)
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("idx %d: got %+v want %+v", i, got[i], w)
		}
	}
}

func TestFetchWatchedRespectsMax(t *testing.T) {
	fixture, _ := os.ReadFile("testdata/events.json")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(fixture)
	}))
	defer srv.Close()

	got, _ := FetchWatched(srv.URL, "x", "", 2)
	if len(got) != 2 {
		t.Errorf("want 2, got %d", len(got))
	}
}

func TestFetchWatchedNon200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message":"rate limit"}`))
	}))
	defer srv.Close()

	_, err := FetchWatched(srv.URL, "x", "", 8)
	if err == nil {
		t.Fatal("expected error on 403")
	}
}
