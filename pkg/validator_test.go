package pkg

import "testing"

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "1_no scheme",
		input:       "google.com",
		expected:    false,
	},
	{
		description: "2_only raw text",
		input:       "google",
		expected:    false,
	},
	{
		description: "3_http url",
		input:       "http://google.com",
		expected:    true,
	},
	{
		description: "4_https url",
		input:       "https://google.com",
		expected:    true,
	},
	{
		description: "5_www url",
		input:       "https://www.google.com",
		expected:    true,
	},
	{
		description: "6_host start with dot",
		input:       "http://.google.com",
		expected:    false,
	},
	{
		description: "7_invalid protocol",
		input:       "ftp://google.com",
		expected:    false,
	},
	{
		description: "8_url with query",
		input:       "https://www.google.com/search?q=golang",
		expected:    true,
	},
	{
		description: "9_url with parameter",
		input:       "http://test.com/api/get?origin=https://www.google.com",
		expected:    true,
	},
	{
		description: "10_url with parameter and ivalid host",
		input:       "http://test/api/get?origin=https://www.google.com",
		expected:    false,
	},
	{
		description: "11_url with path",
		input:       "http://test.com:/api/all",
		expected:    true,
	},
	{
		description: "12_url with port",
		input:       "http://test.com:8080",
		expected:    true,
	},
	{
		description: "13_invalid scheme",
		input:       "http:/google.com",
		expected:    false,
	},
	{
		description: "14_another invalid scheme",
		input:       "http//google.com",
		expected:    false,
	},
}

func TestIsValidURL(t *testing.T) {
	for _, v := range testCases {
		t.Run(v.description, func(t *testing.T) {
			got := IsValidURL(v.input)

			if got != v.expected {
				t.Errorf("got: %v, want: %v\n", got, v.expected)
			}
		})
	}
}
