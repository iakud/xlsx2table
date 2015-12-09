package main

import (
	"github.com/tealeg/xlsx"
	"log"
	"path/filepath"
)

type Parser struct {
	inputpath  string
	outputpath string
	language   string
	/*
		name      string
		sheet     *xlsx.Sheet*/
	keyfields []*Field
	fields    []*Field
}

func NewParser(inputpath, outputpath, language string) *Parser {
	return &Parser{inputpath: inputpath, outputpath: outputpath, language: language}
}

func (this *Parser) ParseTable(table *Table) {
	filename := filepath.Join(this.inputpath, table.File)
	file, err := xlsx.OpenFile(filename)
	if err != nil {
		log.Println(err)
		return
	}

	if sheet, ok := file.Sheet[table.Sheet]; ok {
		if this.checkField(sheet) {
			this.parseFields(sheet, table.Columns)
		}
	} else {
		log.Println(filename, "not found.")
	}
}

func (this *Parser) checkField(sheet *xlsx.Sheet) bool {
	if sheet.MaxRow < 2 {
		return false
	}
	return true
}

func (this *Parser) parseFields(sheet *xlsx.Sheet, columns []Column) bool {
	row := sheet.Rows[0]
	for _, column := range columns {
		if field, ok := this.parseField(row, column); ok {
			if column.Key {
				this.keyfields = append(this.keyfields, field)
			} else {
				this.fields = append(this.fields, field)
			}
		} else {
			return false
		}
	}
	return true
}

func (this *Parser) parseField(row *xlsx.Row, column Column) (*Field, bool) {
	for index, cell := range row.Cells {
		if cell.String() == column.Title {
			return &Field{column.Name, index, column.Type, column.Title}, false
		}
	}
	return nil, false
}

func (this *Parser) export() {

}
