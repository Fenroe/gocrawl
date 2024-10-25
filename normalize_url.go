package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	u, err := url.Parse(inputURL)
	if err != nil {
		return "", errors.New("couldn't parse URL")
	}
	fmt.Println(u.Host)
	newURL := u.Host + u.Path
	newURL = strings.TrimSuffix(newURL, "/")
	if u.RawQuery != "" {
		newURL += fmt.Sprintf("?%s", u.RawQuery)
	}
	if u.Fragment != "" {
		newURL += fmt.Sprintf("#%s", u.Fragment)
	}
	newURL = strings.ToLower(newURL)
	return newURL, nil
}
