package render

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	"github.com/khalldev/readme-profile/internal/github"
)

//go:embed templates/README.tmpl
var tmplFS embed.FS

type Data struct {
	User    string
	Watched []github.Watched
}

func Render(d Data) (string, error) {
	t, err := template.ParseFS(tmplFS, "templates/README.tmpl")
	if err != nil {
		return "", fmt.Errorf("parse template: %w", err)
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, d); err != nil {
		return "", fmt.Errorf("execute template: %w", err)
	}
	return buf.String(), nil
}
