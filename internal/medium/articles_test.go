package medium

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFetchArticles(t *testing.T) {
	fixture, err := os.ReadFile("testdata/feed.xml")
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/rss+xml")
		w.Write(fixture)
	}))
	defer srv.Close()

	got, err := FetchArticles(srv.URL, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 {
		t.Fatalf("want 3 articles, got %d", len(got))
	}
	if got[0].Title != "Type-Driven Design in Swift" {
		t.Errorf("title mismatch: %q", got[0].Title)
	}
	if got[0].Link != "https://medium.com/@khalkhalkhal/type-driven-1" {
		t.Errorf("link mismatch: %q", got[0].Link)
	}
}

func TestFetchArticlesMax(t *testing.T) {
	fixture, _ := os.ReadFile("testdata/feed.xml")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(fixture)
	}))
	defer srv.Close()

	got, _ := FetchArticles(srv.URL, 2)
	if len(got) != 2 {
		t.Errorf("max=2 gave %d", len(got))
	}
}

func TestFetchArticlesNon200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()
	_, err := FetchArticles(srv.URL, 10)
	if err == nil {
		t.Fatal("want error")
	}
}

func TestFeedURL(t *testing.T) {
	if FeedURL("alice") != "https://medium.com/feed/@alice" {
		t.Fail()
	}
}
