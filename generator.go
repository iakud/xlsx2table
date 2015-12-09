package main

import (
	"github.com/tealeg/xlsx"
	"log"
	"path/filepath"
)

type Generator struct {
	InputPath  string
	OutputPath string
	Language   string
	/*
		name      string
		sheet     *xlsx.Sheet*/
	keyfields []*Field
	fields    []*Field
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (this *Generator) ParseTable(table *Table) {
	filename := filepath.Join(this.InputPath, table.File)
	file, err := xlsx.OpenFile(filename)
	if err != nil {
		log.Println(err)
		return
	}

	if sheet, ok := file.Sheet[table.Sheet]; ok {
		//fmt.Errorf("xl/_rels/workbook.xml.rels not found in input xlsx.")
		this.parseFields(sheet, table.Columns)
	}
}

func (this *Generator) checkField(sheet *xlsx.Sheet) bool {
	if sheet.MaxRow < 2 {
		return false
	}
	return true
}

func (this *Generator) parseFields(sheet *xlsx.Sheet, columns []Column) bool {
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

func (this *Generator) parseField(row *xlsx.Row, column Column) (*Field, bool) {
	for index, cell := range row.Cells {
		if cell.String() == column.Title {
			return &Field{column.Name, index, column.Type, column.Title}, false
		}
	}
	return nil, false
}
func (this *Generator) export() {

}
