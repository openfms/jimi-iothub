package utils

import (
	"errors"
	"gotest.tools/v3/assert"
	"net/url"
	"testing"
	"time"
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

func TestAddOffsetToUnixTime(t *testing.T) {
	now := time.Now()
	testCases := map[string]struct {
		offsetStr string
		unixTime  int64
		expected  int64
		errWant   error
	}{
		// Test cases with valid input
		"PositiveOffset": {
			offsetStr: "+02:00",
			unixTime:  now.Unix(),
			expected:  now.Add(time.Hour * 2).Unix(),
			errWant:   nil,
		},
		"NegativeOffset": {
			offsetStr: "-05:00",
			unixTime:  now.Unix(),
			expected:  now.Add(time.Hour * -5).Unix(),
			errWant:   nil,
		},
		// Test cases with invalid input
		"InvalidOffset": {
			offsetStr: "invalid",
			unixTime:  1596211200,
			expected:  0,
			errWant:   errors.New("invalid offset format"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := AddOffsetToUnixTime(test.offsetStr, test.unixTime)

			if test.errWant != nil {
				assert.Equal(t, test.errWant.Error(), err.Error())
			} else {
				assert.NilError(t, err)
				assert.Equal(t, result, test.expected)
			}
		})
	}
}
