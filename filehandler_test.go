package main

import (
	"testing"
)

func TestSeveralResolveFile(t *testing.T) {
	var tests = []struct {
		filepath string
		url      string
		expected string
	}{
		{"test-files/", "/", "test-files/index.json"},
		{"test-files/users/1", "/users/1", "test-files/users/1.json"},
		{"test-files/find-first/match", "/find-first/match/", "test-files/find-first/match2.json"},
		{"test-files/users/2", "/users/2", "test-files/users/404.json"},
		{"test-files/whatever/doesnotexist", "/whatever/doesnotexist", "test-files/404.json"},
	}

	for _, test := range tests {
		parsed, err := resolveFile("test-files", test.filepath, test.url, nil)

		if err != nil {
			t.Error(err)
			return
		}

		if parsed == nil {
			t.Errorf("Parsed file %s was nil, probably because none matched\n", test.filepath)
			return
		}

		if parsed.FilePath != test.expected {
			t.Errorf("Expecting %s but got %s\n", test.expected, parsed.FilePath)
			return
		}
	}
}

//this is an actual integration test, find a way to isolate it from the test-files directory
func TestFindFirst(t *testing.T) {
	parsed, err := findFirst("test-files/find-first", "match", "/find-first/match/", nil)

	if err != nil {
		t.Error(err)
		return
	}

	if parsed == nil {
		t.Error("Parsed file was nil, probably because none matched")
		return
	}

	if parsed.FilePath != "test-files/find-first/match2.json" {
		t.Error("Expecting file test-files/find-first/match2.json but got", parsed.FilePath)
		return
	}
}
