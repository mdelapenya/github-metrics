package github

import (
	"github.com/mdelapenya/github-metrics/http"
)

func IssuesCountByLabel(owner string, repository string, label string) ([]byte, error) {
	return search(owner, repository, "label:"+label)
}

func Labels(owner string, repository string) ([]byte, error) {
	return repoGet("labels", owner, repository)
}

func repoGet(resource string, owner string, repository string) ([]byte, error) {
	req := http.Request{
		URL: "https://api.github.com/repos/" + owner + "/" + repository + "/" + resource,
	}

	return http.Get(req)
}

func search(owner string, repository string, queryString string) ([]byte, error) {
	req := http.Request{
		URL: "https://api.github.com/search/issues?q=repo:" + owner + "/" + repository + "+" + queryString,
		Headers: map[string]string{
			"User-Agent": "request",
		},
	}

	return http.Get(req)
}
