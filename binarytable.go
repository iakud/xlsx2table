package main

import (
	"errors"
	"github.com/tealeg/xlsx"
	"io"
	"os"
	"strconv"
	"strings"
)

type typeWriter struct {
	*binaryWriter
}

func newTypeWriter(w io.Writer) *typeWriter {
	return &typeWriter{newBinaryWriter(w)}
}

func (tw *typeWriter) WriteBoolString(str string) error {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return err
	}
	return tw.WriteBool(b)
}

func (tw *typeWriter) WriteInt8String(str string) error {
	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return err
	}
	return tw.WriteInt8(int8(i))
}

func (tw *typeWriter) WriteInt16String(str string) error {
	i, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return err
	}
	return tw.WriteInt16(int16(i))
}

func (tw *typeWriter) WriteInt32String(str string) error {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return err
	}
	return tw.WriteInt32(int32(i))
}

func (tw *typeWriter) WriteInt64String(str string) error {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	return tw.WriteInt64(int64(i))
}

func (tw *typeWriter) WriteUint8String(str string) error {
	n, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return err
	}
	return tw.WriteUint8(uint8(n))
}

func (tw *typeWriter) WriteUint16String(str string) error {
	n, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return err
	}
	return tw.WriteUint16(uint16(n))
}

func (tw *typeWriter) WriteUint32String(str string) error {
	n, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return err
	}
	return tw.WriteUint32(uint32(n))
}

func (tw *typeWriter) WriteUint64String(str string) error {
	n, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}
	return tw.WriteUint64(uint64(n))
}

func (tw *typeWriter) WriteFloatString(str string) error {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return err
	}
	return tw.WriteFloat(float32(f))
}

func (tw *typeWriter) WriteDoubleString(str string) error {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}
	return tw.WriteDouble(float64(f))
}

func (tw *typeWriter) WriteKind(kind Kind, str string) error {
	str = strings.TrimSpace(str)
	switch kind {
	case Bool:
		return tw.WriteBoolString(str)
	case Int8:
		return tw.WriteInt8String(str)
	case Int16:
		return tw.WriteInt16String(str)
	case Int32:
		return tw.WriteInt32String(str)
	case Int64:
		return tw.WriteInt64String(str)
	case Uint8:
		return tw.WriteUint8String(str)
	case Uint16:
		return tw.WriteUint16String(str)
	case Uint32:
		return tw.WriteUint32String(str)
	case Uint64:
		return tw.WriteUint64String(str)
	case Float:
		return tw.WriteFloatString(str)
	case Double:
		return tw.WriteDoubleString(str)
	case String:
		return tw.WriteString(str)
	}
	return errors.New("Invalid type")
}

func (tw *typeWriter) WriteVector(value Kind, str string) error {
	values := strings.Split(strings.TrimSpace(str), ",")
	if err := tw.WriteUint32(uint32(len(values))); err != nil {
		return err
	}
	for _, valuestr := range values {
		if err := tw.WriteKind(value, valuestr); err != nil {
			return err
		}
	}
	return nil
}

func (tw *typeWriter) WriteMap(key Kind, value Kind, str string) error {
	pairs := strings.Split(strings.TrimSpace(str), ",")
	if err := tw.WriteUint32(uint32(len(pairs))); err != nil {
		return err
	}
	for _, pair := range pairs {
		keyvalue := strings.Split(pair, "=")
		if len(keyvalue) != 2 {
			return errors.New("Invalid map")
		}
		if err := tw.WriteKind(key, keyvalue[0]); err != nil {
			return err
		}
		if err := tw.WriteKind(value, keyvalue[1]); err != nil {
			return err
		}
	}
	return nil
}

func (tw *typeWriter) WriteType(ftype Type, str string) error {
	if ftype.Kind() == Vector {
		return tw.WriteVector(ftype.Value(), str)
	}
	if ftype.Kind() == Map {
		return tw.WriteMap(ftype.Key(), ftype.Value(), str)
	}
	return tw.WriteKind(ftype.Kind(), str)
}

type binaryTable struct {
	keys   []*Field
	fields []*Field
}

func newBinaryTable(keys []*Field, fields []*Field) *binaryTable {
	return &binaryTable{keys, fields}
}

func (bt *binaryTable) WriteFile(filename string, rows []*xlsx.Row) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	tw := newTypeWriter(f)
	tw.WriteUint32(uint32(len(rows)))
	for _, row := range rows {
		for _, field := range bt.keys {
			if err := tw.WriteType(field.Type(), row.Cells[field.Column()].Value); err != nil {
				return err
			}
		}
		for _, field := range bt.fields {
			if err := tw.WriteType(field.Type(), row.Cells[field.Column()].Value); err != nil {
				return err
			}
		}
	}
	return nil
}
