package main

import (
	"encoding/xml"
	"io/ioutil"
)

type TableSet struct {
	InputPath  string  `xml:"inputpath,attr"`
	OutputPath string  `xml:"outputpath,attr"`
	Language   string  `xml:"language,attr"`
	Tables     []Table `xml:"table"`
}

type Table struct {
	Name    string   `xml:"name,attr"`
	File    string   `xml:"file,attr"`
	Sheet   string   `xml:"sheet,attr"`
	Columns []Column `xml:"column"`
}

type Column struct {
	Name  string `xml:"name,attr"`
	Key   bool   `xml:"key,attr"`
	Title string `xml:"title,attr"`
	Type  string `xml:"type,attr"`
}

func (this *TableSet) ParseFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(content, &this)
	if err != nil {
		return err
	}
	return nil
}
