package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mdelapenya/github-metrics/log"
	"go.uber.org/zap"
)

// Request configures an HTTP request
type Request struct {
	BasicAuthUser     string
	BasicAuthPassword string
	EncodeURL         bool
	Headers           map[string]string
	method            string
	Payload           string // string representation of the payload, in JSON format
	QueryString       string
	URL               string
}

// GetURL returns the URL as a string
func (req *Request) GetURL() string {
	if req.QueryString == "" {
		return req.URL
	}

	u := req.URL + "?"
	if req.EncodeURL {
		return u + url.QueryEscape(req.QueryString)
	}

	return u + req.QueryString
}

// Delete executes a DELETE request
func Delete(r Request) ([]byte, error) {
	r.method = "DELETE"

	return request(r)
}

// Head executes a HEAD request
func Head(r Request) ([]byte, error) {
	r.method = "HEAD"

	return request(r)
}

// Get executes a GET request
func Get(r Request) ([]byte, error) {
	r.method = "GET"

	return request(r)
}

// Post executes a POST request
func Post(r Request) ([]byte, error) {
	r.method = "POST"

	return request(r)
}

// Put executes a PUT request
func Put(r Request) ([]byte, error) {
	r.method = "PUT"

	return request(r)
}

// Post executes a request
func request(r Request) ([]byte, error) {
	escapedURL := r.GetURL()

	errorArgsFunc := func(err error) []log.Field {
		return []log.Field{
			zap.Error(err),
			zap.String("method", r.method),
			zap.String("url", escapedURL),
		}
	}

	var body io.Reader
	if r.Payload != "" {
		body = bytes.NewReader([]byte(r.Payload))
		errorArgsFunc = func(err error) []log.Field {
			return []log.Field{
				zap.Error(err),
				zap.String("method", r.method),
				zap.String("payload", r.Payload),
				zap.String("url", escapedURL),
			}
		}
	} else {
		body = nil
	}

	req, err := http.NewRequest(r.method, escapedURL, body)
	if err != nil {
		log.Error("error creating the request", errorArgsFunc(err)...)
		return []byte{}, err
	}

	if r.Headers != nil {
		for k, v := range r.Headers {
			req.Header.Set(k, v)
		}
	}

	if r.BasicAuthUser != "" {
		req.SetBasicAuth(r.BasicAuthUser, r.BasicAuthPassword)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("error performing the request", errorArgsFunc(err)...)
		return []byte{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("error reading the body response", errorArgsFunc(err)...)
		return []byte{}, err
	}

	// http.Status ==> [2xx, 4xx)
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		return bodyBytes, nil
	}

	return bodyBytes, fmt.Errorf("%s request failed with %d", r.method, resp.StatusCode)
}
