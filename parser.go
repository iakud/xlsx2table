package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

var FieldTypeMap = map[string]int{
	"bool":   FieldTypeBool,
	"int8":   FieldTypeInt8,
	"uint8":  FieldTypeUInt8,
	"int16":  FieldTypeInt16,
	"uint16": FieldTypeUInt16,
	"int32":  FieldTypeInt32,
	"uint32": FieldTypeUInt32,
	"int64":  FieldTypeInt64,
	"uint64": FieldTypeUInt64,
	"float":  FieldTypeFloat32,
	"double": FieldTypeFloat64,
	"string": FieldTypeString,
}

func ParseColumn(row *xlsx.Row, title string) (int, error) {
	for column, cell := range row.Cells {
		if cell.String() == title {
			return column, nil
		}
	}
	return 0, fmt.Errorf("unknow column.")
}

func ParserFields(row *xlsx.Row, columns []Column) ([]*Field, []*Field, error) {
	var keyfields, fields []*Field
	fieldmap := make(map[string]*Field)
	for _, column := range columns {
		if _, ok := fieldmap[column.Name]; ok {
			return nil, nil, fmt.Errorf("name repeat.")
		}
		fieldtype, ok := FieldTypeMap[column.Type]
		if !ok {
			return nil, nil, fmt.Errorf("unknow field.")
		}

		fieldcolumn, err := ParseColumn(row, column.Title)
		if err != nil {
			return nil, nil, err
		}

		field := &Field{column.Name, column.Title, fieldtype, fieldcolumn}
		if column.Key {
			keyfields = append(keyfields, field)
		} else {
			fields = append(fields, field)
		}
		fieldmap[field.Name] = field
	}
	return keyfields, fields, nil
}
