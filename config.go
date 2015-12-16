package main

import (
	"encoding/xml"
	"io/ioutil"
)

type tableset struct {
	Tables []table `xml:"table"`
}

type table struct {
	Name    string   `xml:"name,attr"`
	File    string   `xml:"file,attr"`
	Sheet   string   `xml:"sheet,attr"`
	Columns []column `xml:"column"`
}

type column struct {
	Name  string `xml:"name,attr"`
	Key   bool   `xml:"key,attr"`
	Title string `xml:"title,attr"`
	Type  string `xml:"type,attr"`
}

func ParseConfigFile(filename string) ([]table, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var set tableset
	err = xml.Unmarshal(content, &set)
	if err != nil {
		return nil, err
	}
	return set.Tables, nil
}
