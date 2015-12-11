package main

import (
	"flag"
	"log"
	"reflect"
)

var FileName, InputPath, OutputPath, Language string

func main() {
	//var datas [][]interface{}
	//var row []interface{}
	//var m map[int]string
	//t := reflect.TypeOf(m)
	reflect.Kind
	flag.StringVar(&FileName, "f", "", "description file")
	flag.StringVar(&InputPath, "i", "", "input path")
	flag.StringVar(&OutputPath, "o", "", "output path")
	flag.StringVar(&Language, "l", "", "language")
	flag.Parse()

	tableset := NewTableSet()
	err := tableset.Parse(FileName)
	if err != nil {
		log.Fatalln(err)
	}

	for _, table := range tableset.Tables {
		err = ExportTable(&table)
		if err != nil {
			log.Println(err)
		}
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
