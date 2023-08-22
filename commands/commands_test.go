package commands

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestGenerateCommand(t *testing.T) {
	tests := map[string]struct {
		command  DeviceCommand
		params   []string
		expected string
	}{
		"NoParams": {
			command:  COREKITSW,
			params:   nil,
			expected: "COREKITSW",
		},
		"WithParams": {
			command:  RSERVICE,
			params:   []string{"192.168.0.1:1935/live"},
			expected: "RSERVICE,192.168.0.1:1935/live",
		},
		"MultipleParams": {
			command:  UPLOAD,
			params:   []string{"http://www.example.com/upload", "param2", "param3"},
			expected: "UPLOAD,http://www.example.com/upload,param2,param3",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			result := GenerateCommand(test.command, test.params...)
			assert.Equal(t, test.expected, result)
		})
	}
}
