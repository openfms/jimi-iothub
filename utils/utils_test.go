package utils

import (
	"errors"
	"gotest.tools/v3/assert"
	"net/url"
	"testing"
)

func TestGetEndpointURL(t *testing.T) {
	tests := map[string]struct {
		endPoint string
		errWant  error
	}{
		"success domain": {
			errWant:  nil,
			endPoint: "https://example.com",
		},
		"success domain with port": {
			errWant:  nil,
			endPoint: "https://sub.example.com:8080",
		},
		"success ip": {
			errWant:  nil,
			endPoint: "http://192.168.1.1",
		},
		"success ip port": {
			errWant:  nil,
			endPoint: "http://192.168.1.1:8080",
		},
		"invalid": {
			errWant:  errors.New("invalid.endpoint is an invalid endpoint"),
			endPoint: "invalid.endpoint",
		},
		"invalid ip": {
			endPoint: "256.256.256.256",
			errWant:  errors.New("256.256.256.256 is an invalid endpoint"),
		},
		"invalid path": {
			endPoint: "http://192.168.1.1:8080/asdad",
			errWant:  errors.New("http://192.168.1.1:8080/asdad is an invalid endpoint"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			u, err := url.Parse(test.endPoint)
			assert.NilError(t, err)
			err = isValidEndpoint(u)
			if test.errWant != nil {
				assert.ErrorContains(t, err, test.errWant.Error())
			} else {
				assert.NilError(t, err)
			}
		})
	}
}
func Test_generateUniqueInstructionID(t *testing.T) {
	id := GenerateUniqueInstructionID()
	t.Log(id)
}

func TestGenerateVideoLink(t *testing.T) {
	tests := map[string]struct {
		endpoint string
		prefix   string
		channel  int
		secure   bool
		imei     string
		expected string
	}{
		"with prefix": {
			endpoint: "example.com:8881",
			prefix:   "live",
			channel:  0,
			secure:   false,
			imei:     "86271111111111",
			expected: "http://example.com:8881/live/0/86271111111111.flv",
		},
		"without prefix": {
			endpoint: "example.com:8881",
			prefix:   "",
			secure:   false,
			channel:  0,
			imei:     "86271111111111",
			expected: "http://example.com:8881/0/86271111111111.flv",
		},
		"with prefix secure": {
			endpoint: "example.com:8881",
			prefix:   "live",
			channel:  0,
			secure:   true,
			imei:     "86271111111111",
			expected: "https://example.com:8881/live/0/86271111111111.flv",
		},
		"without prefix secure": {
			endpoint: "example.com:8881",
			prefix:   "",
			secure:   true,
			channel:  0,
			imei:     "86271111111111",
			expected: "https://example.com:8881/0/86271111111111.flv",
		},
		// Add more test cases as needed
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := GenerateHttpFLVLink(test.secure, test.endpoint, test.prefix, test.channel, test.imei)
			assert.NilError(t, err)
			assert.Equal(t, test.expected, result, "Unexpected result")
		})
	}
}

func TestGenerateRtmpLink(t *testing.T) {
	tests := map[string]struct {
		endpoint string
		prefix   string
		channel  int
		secure   bool
		imei     string
		expected string
	}{
		"with prefix": {
			endpoint: "example.com:8881",
			prefix:   "live",
			channel:  0,
			secure:   false,
			imei:     "862711111111111",
			expected: "rtmp://example.com:8881/live/0/862711111111111",
		},
		"without prefix": {
			endpoint: "example.com:8881",
			prefix:   "",
			secure:   false,
			channel:  0,
			imei:     "862711111111111",
			expected: "rtmp://example.com:8881/0/862711111111111",
		},
		"with prefix secure": {
			endpoint: "example.com:8881",
			prefix:   "live",
			channel:  0,
			secure:   true,
			imei:     "862711111111111",
			expected: "rtmps://example.com:8881/live/0/862711111111111",
		},
		"without prefix secure": {
			endpoint: "example.com:8881",
			prefix:   "",
			secure:   true,
			channel:  0,
			imei:     "862711111111111",
			expected: "rtmps://example.com:8881/0/862711111111111",
		},
		// Add more test cases as needed
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := GenerateRtmpLink(test.secure, test.endpoint, test.prefix, test.channel, test.imei)
			assert.NilError(t, err)
			assert.Equal(t, test.expected, result, "Unexpected result")
		})
	}
}
