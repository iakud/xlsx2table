package main

import (
	"os"
	"text/template"
)

var templateNames = []string{
	"template/cpp.template",
}

type codetemplate struct {
	t *template.Template
}

func newCodeTemplate() *codetemplate {
	return &codetemplate{template.Must(template.ParseFiles(templateNames...))}
}

func (ct *codetemplate) WriteFile(filename string, fieldtable *FieldTable) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return ct.t.Execute(f, fieldtable)
}
