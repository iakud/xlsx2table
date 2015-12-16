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
	fieldmap := make(map[string]*Field)
	for _, column := range columns {
		if _, ok := fieldmap[column.Name]; ok {
			return nil, errors.New("repeat column.")
		}
		field, err := parseField(row, &column)
		if err != nil {
			return nil, err
		}
		if column.Key {
			keys = append(keys, field)
		} else {
			fields = append(fields, field)
		}
		fieldmap[column.Name] = field
	}
	if len(keys) == 0 {
		return nil, errors.New("cannot find key.")
	}

	return NewFieldTable(name, keys, fields), nil
}
