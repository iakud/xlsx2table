package main

import (
	"strings"
)

type Field struct {
	name   string
	title  string
	ctype  Type
	column int
}

func NewField(name string, title string, ctype Type, column int) *Field {
	return &Field{name, title, ctype, column}
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Title() string {
	return f.title
}

func (f *Field) Type() Type {
	return f.ctype
}

func (f *Field) Column() int {
	return f.column
}

func (f *Field) String() string {
	return f.Type().String() + " " + f.Name() + "; // " + f.Title()
}

type FieldTable struct {
	name   string
	keys   []*Field
	fields []*Field
}

func NewFieldTable(name string, keys []*Field, fields []*Field) *FieldTable {
	return &FieldTable{name, keys, fields}
}

func (ft *FieldTable) ComposedKey() bool {
	return len(ft.keys) > 1
}

func (ft *FieldTable) KeyType() string {
	if ft.ComposedKey() {
		return ft.name + "Id"
	} else {
		return ft.keys[0].Type().String()
	}
}

func (ft *FieldTable) KeyName() string {
	if ft.ComposedKey() {
		return "id"
	} else {
		return ft.keys[0].Name()
	}
}

func (ft *FieldTable) KeyString() string {
	if ft.ComposedKey() {
		return ft.name + "Id" + " " + "id;"
	} else {
		return ft.keys[0].String()
	}
}

func (ft *FieldTable) KeyFields() []*Field {
	return ft.keys
}

func (ft *FieldTable) Fields() []*Field {
	return ft.fields
}

func (ft *FieldTable) Name() string {
	return ft.name
}

func (ft *FieldTable) UpperName() string {
	return strings.ToUpper(ft.Name())
}

func (ft *FieldTable) TableName() string {
	return ft.Name() + "Table"
}

func (ft *FieldTable) ItemName() string {
	return ft.Name() + "Item"
}

func (ft *FieldTable) TableFileName() string {
	return ft.TableName() + ".h"
}

func (ft *FieldTable) BinaryFileName() string {
	return ft.TableName() + ".binary"
}
