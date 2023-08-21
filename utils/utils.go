package utils

import (
	"fmt"
	"net/url"
	"regexp"
)

// GetEndpointURL - construct a new endpoint.
func GetEndpointURL(endpoint string) (*url.URL, error) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	// Validate incoming endpoint URL.
	if e := isValidEndpoint(endpointURL); e != nil {
		return nil, e
	}
	return endpointURL, nil
}

func isValidEndpoint(u *url.URL) error {
	// Regular expression to match domain or IP address with optional port
	// Matches examples: example.com, sub.example.com:8080, 192.168.1.1, 192.168.1.1:8080
	regex := `^(?:(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}|(?:\d{1,3}\.){3}\d{1,3})(?::\d{1,5})?$`

	if !regexp.MustCompile(regex).MatchString(u.Host) || !(u.Scheme == "http" || u.Scheme == "https") || u.Path != "" {
		return fmt.Errorf("%s is an invalid endpoint", u)
	}
	return nil
}
