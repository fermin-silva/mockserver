package main

import (
	"fmt"
	"testing"
)

func TestFindEndingFilename(t *testing.T) {
	// url `/users/1` should match `test-files/users/1.json`
	parsed, err := resolveFile("test-files", "test-files/users/1", "/users/1", nil)

	fmt.Println("err:", err)
	fmt.Println("parsed:", parsed)
}

//this is an actual integration test, find a way to isolate it from the test-files directory
func TestFindFirst(t *testing.T) {
	parsed, err := findFirst("test-files/find-first", "match", "/find-first/match/", nil)

	if err != nil {
		t.Error(err)
	}

	if parsed == nil {
		t.Error("Parsed file was nil, probably because none matched")
	}

	if parsed.FilePath != "test-files/find-first/match2.json" {
		t.Error("Expecting file test-files/find-first/match2.json but got", parsed.FilePath)
	}
}
