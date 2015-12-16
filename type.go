package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int8
	Int16
	Int32
	Int64
	Uint8
	Uint16
	Uint32
	Uint64
	Float
	Double
	String
	Vector
	Map
)

func (k Kind) String() string {
	if int(k) < len(kindNames) {
		return kindNames[k]
	}
	return "kind" + strconv.Itoa(int(k))
}

var kindNames = []string{
	Invalid: "invalid",
	Bool:    "bool",
	Int8:    "int8_t",
	Int16:   "int16_t",
	Int32:   "int32_t",
	Int64:   "int64_t",
	Uint8:   "uint8_t",
	Uint16:  "uint16_t",
	Uint32:  "uint32_t",
	Uint64:  "uint64_t",
	Float:   "float",
	Double:  "double",
	String:  "std::string",
}

type Type interface {
	Kind() Kind
	Key() Kind
	Value() Kind
	String() string
}

type ctype struct {
	kind Kind
}

func (t *ctype) Kind() Kind {
	return t.kind
}

type vectorType struct {
	ctype
	value Kind
}

type mapType struct {
	ctype
	key   Kind
	value Kind
}

func (t *ctype) Key() Kind {
	if t.Kind() != Map {
		panic("ctype: Key of invalid type")
	}
	return (*mapType)(unsafe.Pointer(t)).key
}

func (t *ctype) Value() Kind {
	switch t.Kind() {
	case Vector:
		return (*vectorType)(unsafe.Pointer(t)).value
	case Map:
		return (*mapType)(unsafe.Pointer(t)).value
	}
	panic("ctype: Value of invalid type")
}

func (t *ctype) String() string {
	return t.Kind().String()
}

func (t *vectorType) String() string {
	return "std::vector<" + t.value.String() + ">"
}

func (t *mapType) String() string {
	return "std::map<" + t.key.String() + "," + t.value.String() + ">"
}

func KindOf(str string) Kind {
	switch strings.ToLower(strings.TrimSpace(str)) {
	case "bool":
		return Bool
	case "int8":
		return Int8
	case "int16":
		return Int16
	case "int32":
		return Int32
	case "int64":
		return Int64
	case "uint8":
		return Uint8
	case "uint16":
		return Uint16
	case "uint32":
		return Uint32
	case "uint64":
		return Uint64
	case "float":
		return Float
	case "double":
		return Double
	case "string":
		return String
	}
	return Invalid
}

var vectorRegexp *regexp.Regexp = regexp.MustCompile(`^[\s]*vector[\s]*\[(.*)\][\s]*$`)
var mapRegexp *regexp.Regexp = regexp.MustCompile(`^[\s]*map[\s]*\[(.*),(.*)\][\s]*$`)

func TypeOf(str string) (Type, error) {
	if kind := KindOf(str); kind != Invalid {
		return &ctype{kind}, nil
	}
	if vectorstr := vectorRegexp.FindStringSubmatch(str); len(vectorstr) > 1 {
		return VectorTypeOf(vectorstr[1])
	}
	if mapstr := mapRegexp.FindStringSubmatch(str); len(mapstr) > 2 {
		return MapTypeOf(mapstr[1], mapstr[2])
	}
	return nil, errors.New("Invalid type.")
}

func VectorTypeOf(valuestr string) (Type, error) {
	if value := KindOf(valuestr); value != Invalid {
		return &vectorType{ctype{Vector}, value}, nil
	}
	return nil, errors.New("Invalid type.")
}

func MapTypeOf(keystr string, valuestr string) (Type, error) {
	key := KindOf(keystr)
	value := KindOf(valuestr)
	if key != Invalid && value != Invalid {
		return &mapType{ctype{Map}, key, value}, nil
	}
	return nil, errors.New("Invalid type.")
}
