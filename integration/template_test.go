package main

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	expected := `
{
	"file" : "template.json",
	"path" : "files/template.json"
}
	`

	if err := EqualJsonGet("/template", expected); err != nil {
		t.Error(err)
		return
	}
}

func TestTemplateComplex(t *testing.T) {
	expected := `
{
	"file" : "template_complex.json",
	"path" : "files/template_complex.json",
	"names" : [ "Juan", "Pedro", "Miguel" ],
	"title" : "Title from config",
	"from_include" : true
}
	`

	if err := EqualJsonGet("/template_complex", expected); err != nil {
		t.Error(err)
		return
	}
}
