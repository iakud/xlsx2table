package main

import ()

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
	FieldTypeFloat32
	FieldTypeFloat64
	FieldTypeString
)

type Field struct {
	Name   string
	Title  string
	Type   int
	Column int
}

/*
import (
	"github.com/tealeg/xlsx"
	"strconv"
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
	FieldTypeFloat32
	FieldTypeFloat64
	FieldTypeString
)

type Field struct {
	Name   string
	Title  string
	Type   int
	Column int
}

func NewField(name string, title string, tp int) *Field {
	return &Field{Name: name, Title: title, Type: tp}
}

func (this *Field) Parse(row *xlsx.Row) bool {
	for column, cell := range row.Cells {
		if cell.String() == this.Title {
			this.Column = column
			return true
		}
	}
	return false
}

func (this *Field) ParseData(str string) (interface{}, error) {
	switch tp {
	case FieldTypeBool:
		value, err := strconv.ParseBool(str)
		return value, err
	case FieldTypeInt8:
		v, err := cell.Int()
		if err != nil {
			return nil, err
		}
		return int8(v), nil
	case FieldTypeUInt8:
		v, err := cell.Int()
		if err != nil {
			return nil, err
		}
		return uint8(v), nil
	case FieldTypeInt16:
		v, err := cell.Int()
		if err != nil {
			return nil, err
		}
		return int16(v), nil
	case FieldTypeUInt16:
		v, err := cell.Int()
		if err != nil {
			return nil, err
		}
		return uint16(v), nil
	case FieldTypeInt32:
		v, err := cell.Int()
		if err != nil {
			return nil, err
		}
		return int32(v), nil
	case FieldTypeUInt32:
		v, err := cell.Int()
		if err != nil {
			return nil, err
		}
		return uint32(v), nil
	case FieldTypeInt64:
		v, err := cell.Int64()
		if err != nil {
			return nil, err
		}
		return int64(v), nil
	case FieldTypeUInt64:
		v, err := cell.Int()
		if err != nil {
			return nil, err
		}
		return uint64(v), nil
	case FieldTypeFloat32:
		value, err := atof32(str)
		return value, err
	case FieldTypeFloat64:
		value, err := atof64(str)
		return value, err
	case FieldTypeString:
		return cell.String(), nil
	}
	return nil, fmt.Errorf("Unknow type")
}

func (this *Field) ParseBool(str string) bool {

}
*/
