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
