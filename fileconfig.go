package main

import (
	"github.com/BurntSushi/toml"
	"regexp"
)

var EMPTY_HEADERS = map[string]string{}

type FileConfig struct {
	Headers    map[string]string
	Match      []string
	Template   bool
	Path       string
	genericCfg map[string]interface{} `toml:"-"`
}

func NewFileConfig(filepath, content string) (FileConfig, error) {
	cfg := FileConfig{Path: filepath}

	if _, err := toml.Decode(content, &cfg); err != nil {
		return cfg, err
	}

	//needed so users can use whatever key they want in the templates
	var genericCfg map[string]interface{}
	_, err := toml.Decode(content, &genericCfg)

	cfg.genericCfg = genericCfg

	return cfg, err
}

func (c *FileConfig) Get(name string) interface{} {
	if val, ok := c.genericCfg[name]; ok {
		return val
	}

	return ""
}

func (f *FileConfig) GetHeaders() map[string]string {
	if f.Headers == nil {
		return EMPTY_HEADERS
	}
	return f.Headers
}

func (f *FileConfig) UseTemplate() bool {
	return f.Template
}

func (f *FileConfig) Matches(url string) bool {
	//always true unless there are configured regexes
	if len(f.Match) == 0 {
		return true
	}

	//TODO cache regex compilation
	for _, str := range f.Match {
		rx, _ := regexp.Compile(str)

		if rx.MatchString(url) {
			return true
		}
	}

	return false
}
