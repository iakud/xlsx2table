package main

import (
	"errors"
	"flag"
	"github.com/tealeg/xlsx"
	"log"
	"path/filepath"
)

func main() {
	var filename, inputpath, outputpath string
	flag.StringVar(&filename, "f", "", "description file")
	flag.StringVar(&inputpath, "i", "", "input path")
	flag.StringVar(&outputpath, "o", "", "output path")
	flag.Parse()

	tables, err := ParseConfigFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	codetemplate := newCodeTemplate()

	for _, table := range tables {
		log.Println("export", table.Name)
		sheet, err := openXlsSheet(filepath.Join(inputpath, table.File), table.Sheet)
		if err != nil {
			log.Println(err)
			continue
		}
		fieldtable, err := parseFieldTable(table.Name, sheet.Rows[0], table.Columns)
		if err != nil {
			log.Println(err)
			continue
		}
		binarytable := newBinaryTable(fieldtable.KeyFields(), fieldtable.Fields())
		tablepath := filepath.Join(outputpath, fieldtable.BinaryFileName())
		err = binarytable.WriteFile(tablepath, sheet.Rows[1:])
		if err != nil {
			log.Println(err)
			continue
		}

		codepath := filepath.Join(outputpath, fieldtable.TableFileName())
		err = codetemplate.WriteFile(codepath, fieldtable)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func openXlsSheet(filename string, sheetname string) (*xlsx.Sheet, error) {
	file, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	sheet, ok := file.Sheet[sheetname]
	if !ok {
		return nil, errors.New("sheet not found.")
	}
	if sheet.MaxRow == 0 {
		return nil, errors.New("sheet is empty.")
	}
	return sheet, nil
}
