package github

import (
	"os"
	"strings"

	"github.com/mdelapenya/github-metrics/http"
)

func CsSearch(query string) ([]byte, error) {
	cookie := os.Getenv("GHM_GITHUB_COOKIE")
	q := strings.ReplaceAll(query, "/", "%2F")
	q = strings.ReplaceAll(q, " ", "+")

	req := http.Request{
		URL: "https://cs.github.com/api/count?q=" + q + "&p=1",
		Headers: map[string]string{
			"User-Agent": "request",
			"cookie":     cookie,
		},
	}

	return searchWithBackoff(req)
}
