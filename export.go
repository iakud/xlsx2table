package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func ExportTable(table *Table) error {
	tablepath := filepath.Join(InputPath, table.File)
	file, err := xlsx.OpenFile(tablepath)
	if err != nil {
		return err
	}

	sheet, ok := file.Sheet[table.Sheet]
	if !ok {
		return fmt.Errorf("sheet not found.")
	}

	if sheet.MaxRow == 0 {
		return fmt.Errorf("sheet is empty.")
	}
	keyfields, fields, err := ParserFields(sheet.Rows[0], table.Columns)
	if err != nil {
		return err
	}

	datapath := filepath.Join(OutputPath, table.Name) + ".dat"
	f, err := os.Create(datapath)
	if err != nil {
		return err
	}
	defer f.Close()

	//datatable := NewDataTable(keyfields, fields)
	//datatable.ParseFields(sheet)
	// do export

	//buffer := bytes.NewBuffer(nil)
	WriteTable(f, sheet.Rows[1:], keyfields, fields)
	return nil
}

func WriteTable(w io.Writer, rows []*xlsx.Row, keyfields []*Field, fields []*Field) {
	binary.Write(w, binary.BigEndian, uint32(len(rows)))
	for _, row := range rows {
		buffer := bytes.NewBuffer(nil)
		for _, field := range keyfields {
			WriteField(buffer, row.Cells[field.Column].Value, field.Type)
		}
		for _, field := range fields {
			WriteField(buffer, row.Cells[field.Column].Value, field.Type)
		}
		//length := buffer.Len()
		//binary.Write(w, binary.BigEndian, uint32(length))
		buffer.WriteTo(w)
	}
}

func WriteField(w io.Writer, str string, fieldtype int) {
	fmt.Println(str)
	switch fieldtype {
	case FieldTypeBool:
		value, _ := strconv.ParseBool(str)
		if value {
			binary.Write(w, binary.BigEndian, int8(1))
		} else {
			binary.Write(w, binary.BigEndian, int8(0))
		}
	case FieldTypeInt8:
		value, _ := strconv.ParseInt(str, 10, 8)
		binary.Write(w, binary.BigEndian, int8(value))
	case FieldTypeUInt8:
		value, _ := strconv.ParseUint(str, 10, 8)
		binary.Write(w, binary.BigEndian, uint8(value))
	case FieldTypeInt16:
		value, _ := strconv.ParseInt(str, 10, 16)
		binary.Write(w, binary.BigEndian, int16(value))
	case FieldTypeUInt16:
		value, _ := strconv.ParseUint(str, 10, 16)
		binary.Write(w, binary.BigEndian, uint16(value))
	case FieldTypeInt32:
		value, _ := strconv.ParseInt(str, 10, 32)
		binary.Write(w, binary.BigEndian, int32(value))
	case FieldTypeUInt32:
		value, _ := strconv.ParseUint(str, 10, 32)
		binary.Write(w, binary.BigEndian, uint32(value))
	case FieldTypeInt64:
		value, _ := strconv.ParseInt(str, 10, 64)
		binary.Write(w, binary.BigEndian, int64(value))
	case FieldTypeUInt64:
		value, _ := strconv.ParseUint(str, 10, 64)
		binary.Write(w, binary.BigEndian, uint64(value))
	case FieldTypeFloat32:
		value, _ := strconv.ParseFloat(str, 32)
		binary.Write(w, binary.BigEndian, math.Float32bits(float32(value)))
	case FieldTypeFloat64:
		value, _ := strconv.ParseFloat(str, 64)
		binary.Write(w, binary.BigEndian, math.Float64bits(float64(value)))
	case FieldTypeString:
		length := len(str)
		binary.Write(w, binary.BigEndian, uint32(length))
		io.WriteString(w, str)
	}
}
