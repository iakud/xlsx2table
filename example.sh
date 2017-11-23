#!/bin/bash

config=example/table.xml
inputpath=example/xlsx
outputpath=example
./xlsx2table -f $config -i $inputpath -o $outputpath
