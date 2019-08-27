package main

import (
	"testing"
)

func TestMatches(t *testing.T) {
	var tests = []struct {
		url      string
		expected string
	}{
		{"/matching/", "files/matching/index.json"},
		{"/matching/noexist/", "files/matching/404.json"},
		{"/matching/user/1", "files/matching/user/1.json"},
		{"/matching/user/1234", "files/matching/user/404.json"},
		{"/matching/user/404", "files/matching/user/404.json"},

		//should not match index.json because matching is configured
		{"/matching/users/", "files/matching/users/index_2.json"},

		//with correct file it should match
		{"/matching/users/index", "files/matching/users/index.json"},

		//exact file match should not grab the match config
		{"/matching/users/index.json", "files/matching/users/index.json"},

		//notice how it goes to top level 404 file, and not to files/matching/404.json
		{"/matching/users/noexist/", "files/404.json"},
	}

	for _, test := range tests {
		resp, err := Get(test.url)

		if err != nil {
			t.Error(err)
			return
		}

		if err := resp.HasMatchedFile(test.expected); err != nil {
			t.Error(err)
			return
		}
	}
}
