package main

import (
	//"encoding/binary"
	//	"fmt"
	//	"encoding/binary"
	//	"bytes"
	"github.com/tealeg/xlsx"
	//	"io"
)

type DataTable struct {
	KeyFields []*Field
	Fields    []*Field
}

func NewDataTable(keyfields []*Field, fields []*Field) *DataTable {
	return &DataTable{KeyFields: keyfields, Fields: fields}
}

func (this *DataTable) ParseFields(row *xlsx.Row) {

}

/*
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
*/
/*
func (this *DataTable) WriteTo(w io.Writer) error {
	for _, datarow := range this.DataRows {
		for _, data := range datarow.Datas {
			buffer := bytes.NewBuffer(nil)
			err := this.WriteDataTo(buffer, data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}*/

/*
func (this *DataTable) WriteDataTo(w io.Writer, data interface{}) error {
	switch v := data.(type) {
	case bool:
	case int32:
		binary.Write(w, binary.BigEndian, v)
	case int64:
		binary.Write(w, binary.BigEndian, v)
	case float64:
	case string:
		io.WriteString(w, v)
	default:
		return fmt.Errorf("Unknow type:", v)
	}
	return nil
}
*/
