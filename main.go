package main

import (
	//"encoding/binary"
	"flag"
	"log"
	"path/filepath"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "table.xml", "input description file")
	flag.Parse()

	var tableset TableSet
	err := tableset.ParseFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	relativedir := filepath.Dir(filename)

	inputpath := filepath.Join(relativedir, tableset.InputPath)
	outputpath := filepath.Join(relativedir, tableset.OutputPath)
	parser := NewParser(inputpath, outputpath, tableset.Language)

	for _, table := range tableset.Tables {
		parser.ParseTable(&table)
	}
}

/*
	f, err := os.Create("test.t")
	defer f.Close()
	var tt uint8 = 10
	var ttt uint32 = 55
	var str string = "hello world!"
	binary.Write(f, binary.BigEndian, tt)
	binary.Write(f, binary.BigEndian, ttt)
	binary.Write(f, binary.BigEndian, &str)*/
