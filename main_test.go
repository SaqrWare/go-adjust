package main

import (
	"fmt"
	"testing"
)

// Test FixURL
func TestTableFixURL(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"example.com", "http://example.com"},
		{"http://example.com", "http://example.com"},
		{"https://example.com", "https://example.com"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v, %v", test.input, test.expected), func(t *testing.T) {
			if output := FixURL(test.input); output != test.expected {
				t.Errorf("Test failed: %v inputted, %v expected, %v received", test.input, test.expected, output)
			}
		})

	}
}

func TestCallHTTP(t *testing.T) {
	t.Run("getting_content", func(t *testing.T) {
		url := "https://urlecho.appspot.com/echo?body=test"
		body := CallHTTP(url)
		bodyString := string(body)
		expected := "test"

		if bodyString != expected {
			t.Errorf("Test failed: %v inputted, %v expected, %v received", url, "test", bodyString)
		}
	})
}

func TestHashContent(t *testing.T) {
	t.Run("hash_content", func(t *testing.T) {
		content := []byte("test")
		expected := "098f6bcd4621d373cade4e832627b4f6"
		hash := HashContent(content)

		if hash != expected {
			t.Errorf("Test failed: %v inputted, %v expected, %v received", content, expected, hash)
		}
	})
}

func TestFetchAndHash(t *testing.T) {
	t.Run("get_hashed_content", func(t *testing.T) {
		url := "https://urlecho.appspot.com/echo?body=test"
		expected := "098f6bcd4621d373cade4e832627b4f6"
		result := FetchAndHash(url)
		if result != expected {
			t.Errorf("Test failed: %v inputted, %v expected, %v received", url, expected, result)
		}
	})
}
