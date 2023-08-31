package utils

import (
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"time"
)

const (
	validEndpointRegex = `^(?:(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}|(?:\d{1,3}\.){3}\d{1,3})(?::\d{1,5})?$`
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
	if !regexp.MustCompile(validEndpointRegex).MatchString(u.Host) || !(u.Scheme == "http" || u.Scheme == "https") || u.Path != "" {
		return fmt.Errorf("%s is an invalid endpoint", u)
	}
	return nil
}

func GenerateUniqueInstructionID() string {
	// Generate a random number between 0 and 999999999

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rnd.Intn(1000000000)

	return fmt.Sprintf("%09d", randomNumber)
}
