package render

import (
	"os"
	"testing"

	"github.com/khalldev/readme-profile/internal/github"
)

func TestRenderGolden(t *testing.T) {
	d := Data{
		User: "khalldev",
		Watched: []github.Watched{
			{Name: "BenEmdon/CenteredCollectionView", URL: "https://github.com/BenEmdon/CenteredCollectionView", Date: "2026-05-25"},
			{Name: "kageroumado/phosphene", URL: "https://github.com/kageroumado/phosphene", Date: "2026-05-22"},
		},
	}
	got, err := Render(d)
	if err != nil {
		t.Fatal(err)
	}

	want, err := os.ReadFile("testdata/expected.md")
	if err != nil {
		t.Fatal(err)
	}
	if got != string(want) {
		t.Errorf("output mismatch.\n--- GOT ---\n%s\n--- WANT ---\n%s", got, string(want))
	}
}
