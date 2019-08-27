package main

import (
	"testing"
)

func TestIndex(t *testing.T) {
	expected := `
{
	"file" : "index.json",
	"path" : "{{ File.Path }}"
}
	`

	//notice the {{ File.Path }} is not expanded as this file has no Front Matter

	resp, err := Get("/")

	if err != nil {
		t.Error(err)
		return
	}

	body, _ := resp.Body()

	if ok, err := AreEqualJSON(expected, body); err != nil {
		t.Error(err)
		return
	} else if !ok {
		t.Error("Expecting\n", expected, "\nbut got\n", body)
	}
}

func TestHeaders(t *testing.T) {
	resp, err := Get("/index_headers")

	expected := make(map[string][]string)
	expected["Content-Type"] = []string{"Application/json"}
	expected["Whatever"] = []string{"You Want"}

	//notice how the header is normalized, and thus the Int uppercase letter is
	//lowercased automatically
	expected["Extraint"] = []string{"1234"}

	if err != nil {
		t.Error(err)
		return
	}

	if err := MapContainsExpected(expected, resp.Headers()); err != nil {
		t.Error(err)
		return
	}

	expectedBody := `
{
	"file" : "index_headers.json",
	"path" : "files/index_headers.json"
}
	`

	body, _ := resp.Body()

	if ok, err := AreEqualJSON(expectedBody, body); err != nil {
		t.Error(err)
		return
	} else if !ok {
		t.Error("Expecting\n", expectedBody, "\nbut got\n", body)
	}
}
