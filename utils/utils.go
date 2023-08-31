package utils

import (
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	HttpFLVLinkFormat  = "{{FlvAddress}}/{{Prefix}}/{{channel}}/{{IMEI}}.flv"
	RtmpLinkFormat     = "{{RtmpAddress}}/{{Prefix}}/{{channel}}/{{IMEI}}"
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

// GenerateHttpFLVLink generates flv video link
// Example: http://example.com:8881/live/0/86271111111111.flv
func GenerateHttpFLVLink(secure bool, endpoint, prefix string, channel int, imei string) (string, error) {
	scheme := "http"
	if secure {
		scheme = "https"
	}
	// Construct a secured endpoint URL.
	endpointURLStr := scheme + "://" + endpoint
	endpointURL, err := url.Parse(endpointURLStr)
	if err != nil {
		return "", err
	}
	// Validate incoming endpoint URL.
	if !regexp.MustCompile(validEndpointRegex).MatchString(endpointURL.Host) || endpointURL.Path != "" {
		return "", fmt.Errorf("%s is an invalid endpoint", endpointURL.String())
	}
	replacer := strings.NewReplacer(
		"{{FlvAddress}}", endpointURL.String(),
		"{{Prefix}}", prefix,
		"{{channel}}", fmt.Sprintf("%d", channel),
		"{{IMEI}}", imei,
	)
	format := HttpFLVLinkFormat
	if prefix == "" {
		format = strings.ReplaceAll(format, "{{Prefix}}/", "")
	}
	link := replacer.Replace(format)
	return link, nil
}

// GenerateRtmpLink generates rtmp video link
// Example: rtmp://example.com:1936/live/0/86271111111111
func GenerateRtmpLink(secure bool, endpoint, prefix string, channel int, imei string) (string, error) {
	scheme := "rtmp"
	if secure {
		scheme = "rtmps"
	}
	// Construct a secured endpoint URL.
	endpointURLStr := scheme + "://" + endpoint
	endpointURL, err := url.Parse(endpointURLStr)
	if err != nil {
		return "", err
	}
	// Validate incoming endpoint URL.
	if !regexp.MustCompile(validEndpointRegex).MatchString(endpointURL.Host) || endpointURL.Path != "" {
		return "", fmt.Errorf("%s is an invalid endpoint", endpointURL.String())
	}
	replacer := strings.NewReplacer(
		"{{RtmpAddress}}", endpointURL.String(),
		"{{Prefix}}", prefix,
		"{{channel}}", fmt.Sprintf("%d", channel),
		"{{IMEI}}", imei,
	)
	format := RtmpLinkFormat
	if prefix == "" {
		format = strings.ReplaceAll(format, "{{Prefix}}/", "")
	}
	link := replacer.Replace(format)
	return link, nil
}

func GenerateUniqueInstructionID() string {
	// Generate a random number between 0 and 999999999

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rnd.Intn(1000000000)

	return fmt.Sprintf("%09d", randomNumber)
}
