package main

import (
	"testing"
)

func TestIndex(t *testing.T) {
	expected := `
	{
		"file" : "index.json"
	}
	`

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
