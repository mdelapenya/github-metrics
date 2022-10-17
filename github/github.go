package github

import (
	"fmt"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/mdelapenya/github-metrics/http"
)

func IssuesCountByLabel(owner string, repository string, label string) ([]byte, error) {
	return search(owner, repository, "label:"+label)
}

func Labels(owner string, repository string) ([]byte, error) {
	return repoGet("labels", owner, repository)
}

// getExponentialBackOff returns a preconfigured exponential backoff instance
func getExponentialBackOff(elapsedTime time.Duration) *backoff.ExponentialBackOff {
	var (
		initialInterval     = 500 * time.Millisecond
		randomizationFactor = 0.5
		multiplier          = 2.0
		maxInterval         = 5 * time.Second
		maxElapsedTime      = elapsedTime
	)

	exp := backoff.NewExponentialBackOff()
	exp.InitialInterval = initialInterval
	exp.RandomizationFactor = randomizationFactor
	exp.Multiplier = multiplier
	exp.MaxInterval = maxInterval
	exp.MaxElapsedTime = maxElapsedTime

	return exp
}

func repoGet(resource string, owner string, repository string) ([]byte, error) {
	req := http.Request{
		URL: "https://api.github.com/repos/" + owner + "/" + repository + "/" + resource,
	}

	return searchWithBackoff(req)
}

func search(owner string, repository string, queryString string) ([]byte, error) {
	req := http.Request{
		URL: "https://api.github.com/search/issues?q=repo:" + owner + "/" + repository + "+" + queryString,
		Headers: map[string]string{
			"User-Agent": "request",
		},
	}

	return searchWithBackoff(req)
}

func searchWithBackoff(req http.Request) ([]byte, error) {
	exp := getExponentialBackOff(time.Minute * 1)

	var bytes []byte
	retryCount := 0
	searchFn := func() error {
		b, err := http.Get(req)
		if err != nil {
			retryCount++
			fmt.Printf(">>> [%d][%s] error while processing the Github search for \"%s\". Retrying: %v\n", retryCount, exp.GetElapsedTime(), req.URL, err)
			return err
		}

		bytes = b
		return nil
	}

	err := backoff.Retry(searchFn, exp)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
