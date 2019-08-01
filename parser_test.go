package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	configContent := `
match = [
	"/test/[0-9]+/all"
]

template = true

[Headers]
Content-Type = "Application/json"
Whatever = "Tu Morro"`

	bodyContent := `{
	"hola" : "mundo"
}`

	content := "---\n" + configContent + "\n---\n" + bodyContent

	config, content, err := parse(content)

	assertEqual(t, config, configContent, "Config unexpected")
	fmt.Println("Config:")
	fmt.Println(config)

	assertEqual(t, content, bodyContent, "Body content unexpected")
	fmt.Println("Content:")
	fmt.Println(content)

	if err != nil {
		t.Errorf("Parsing got error %s", err)
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
