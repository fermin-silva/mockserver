package main

import (
	"fmt"
	"github.com/flosch/pongo2"
)

type ParsedFile struct {
	FilePath   string
	content    string
	fileconfig *FileConfig
	config     *Config
}

//TODO add headers, etc
type RequestData interface {
	Query(key string) string
	QueryArray(key string) []string
}

func NewParsedFile(filepath, content string, fileconfig *FileConfig, config *Config) *ParsedFile {
	//TODO check and validate not nulls

	return &ParsedFile{filepath, content, fileconfig, config}
}

func (f *ParsedFile) Matches(url string) bool {
	fmt.Println("checking if", url, "matches")

	if f.fileconfig == nil {
		return true
	}

	return f.fileconfig.Matches(url)
}

func (f *ParsedFile) GetHeaders() map[string]string {
	if f.fileconfig == nil {
		return nil
	}

	return f.fileconfig.GetHeaders()
}

//TODO pass something like "http context", a slimmed down version of a full blown gin context
func (f *ParsedFile) String(rd RequestData) (string, error) {
	if f.fileconfig == nil || !f.fileconfig.UseTemplate() {
		return f.content, nil
	}

	tpl, err := pongo2.FromString(f.content)

	if err != nil {
		return "", err
	}

	ctx := pongo2.Context{}
	ctx["Request"] = rd
	ctx["Config"] = f.config
	ctx["File"] = f.fileconfig

	//TODO add random object to context, to generate strings, numbers, bools, etc

	return tpl.Execute(ctx)
}

func guessContentType() string {
	//TODO based on extension?
	return ""
}
