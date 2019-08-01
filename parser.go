package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ParseFile(filepath string, appconfig *Config) (*ParsedFile, error) {
	bs, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, fmt.Errorf("error while reading file %s: %s", filepath, err)
	}

	config, content, err := parse(string(bs))

	if err != nil {
		return nil, fmt.Errorf("error while parsing file %s: %s", filepath, err)
	}

	var fileconfig FileConfig

	if config != "" {
		fileconfig, err = NewFileConfig(filepath, config)

		if err != nil {
			return nil, fmt.Errorf("error while parsing file config %s. \nContent: %s", err, config)
		}
	}

	return NewParsedFile(filepath, content, fileconfig, appconfig), nil
}

//returns config, file content and error
func parse(content string) (string, string, error) {
	if !strings.HasPrefix(content, "---") {
		return "", content, nil
	}

	firstDashFound := false
	dashStartPosition := -1
	dashCount := 0

	for pos, char := range content {
		if pos < 3 {
			continue
		}

		if char == '-' {
			dashCount++

			if !firstDashFound {
				firstDashFound = true
				dashStartPosition = pos
			}

			if dashCount == 3 {
				return strings.TrimSpace(content[3:dashStartPosition]),
					strings.TrimSpace(content[dashStartPosition+3:]),
					nil
			}
		} else {
			if firstDashFound { //no - char but we were counting --> reset everything
				firstDashFound = false
				dashCount = 0
				dashStartPosition = -1
			}
		}
	}

	return "", "", fmt.Errorf("Found starting '---' but couldn't find the ending '---'")
}
