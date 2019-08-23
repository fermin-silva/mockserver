package main

import (
	"fmt"
	"testing"
)

func TestFileConfig(t *testing.T) {
	configContent := `
match = [
	"/test/[0-9]+/all"
]

template = true

[Headers]
Content-Type = "Application/json"
Whatever = "Tu Morro"
`

	cfg, err := NewFileConfig("filepath", configContent)

	assertEqual(t, cfg.Template, true, "Template should be true")

	headers := cfg.GetHeaders()

	assertEqual(t, headers["Content-Type"], "Application/json", "Content type should be equal")
	assertEqual(t, headers["Whatever"], "Tu Morro", "Whatever header should be equal")

	assertEqual(t, cfg.Matches("/test/12345/all"), true, "Should match given url")

	fmt.Println(cfg, err)
}
