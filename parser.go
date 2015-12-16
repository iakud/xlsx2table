package main

import (
	"errors"
	"github.com/tealeg/xlsx"
)

func parseField(row *xlsx.Row, column *column) (*Field, error) {
	for index, cell := range row.Cells {
		if cell.String() == column.Title {
			ctype, err := TypeOf(column.Type)
			if err != nil {
				return nil, err
			}
			return NewField(column.Name, column.Title, ctype, index), nil
		}
	}
	return nil, errors.New("Unknow column.")
}

func parseFieldTable(name string, row *xlsx.Row, columns []column) (*FieldTable, error) {
	var keys, fields []*Field
	// FIXME :fieldmap := make(map[string]*Field)
	for _, column := range columns {
		field, err := parseField(row, &column)
		if err != nil {
			return nil, err
		}
		if column.Key {
			keys = append(keys, field)
		} else {
			fields = append(fields, field)
		}
	}
	return NewFieldTable(name, keys, fields), nil
}
