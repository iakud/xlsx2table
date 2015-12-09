package main

import (
	"log"
	//	"encoding/binary"
	"github.com/tealeg/xlsx"
	"io"
)

const (
	FieldTypeUnknow = iota
	FieldTypeBool
	FieldTypeInt8
	FieldTypeUInt8
	FieldTypeInt16
	FieldTypeUInt16
	FieldTypeInt32
	FieldTypeUInt32
	FieldTypeInt64
	FieldTypeUInt64
	FieldTypeString
)

type Field struct {
	Name   string
	Column int
	Type   string
	Title  string
}

type Data struct {
	data interface{}
}

type DataRow struct {
	Datas []Data
}

type DataTable struct {
	Fields []Field
	Rows   []DataRow
}

func NewDataTable() *DataTable {
	return &DataTable{}
}

func (this *DataTable) Parse(sheet *xlsx.Sheet) bool {
	return true
}

func (this *DataTable) Write(w io.Writer) error {
	for _ = range this.Rows {

	}
	return nil
}

func (this *DataRow) Write(w io.Writer) error {
	for _ = range this.Datas {

	}
	return nil
}

func (this *Data) Write(w io.Writer) error {
	switch data := this.data.(type) {
	default:

		log.Println("default", data)
	}
	return nil
}
