package parser

import (
	"testing"
)

func Test_ShouldParseLogLine(t *testing.T) {
	testcases := []struct {
		desc              string
		input             string
		expectedIPAddress string
		expectedURL       string
	}{
		{
			"validate parser",
			`192.0.0.0 admin admin [01/Jan/1990:22:22:22 +0100] "GET /test/ HTTP/1.1"`,
			"192.0.0.0",
			"/test/",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			result := parseLogLineContent(testcase.input)

			if result.IPAddress != testcase.expectedIPAddress {
				t.Errorf("Expected IP Address: %s, Got: %s", testcase.expectedIPAddress, result.IPAddress)
			}

			if result.URLPath != testcase.expectedURL {
				t.Errorf("Expected URL path: %s, Got: %s", testcase.expectedURL, result.URLPath)
			}
		})
	}
}
