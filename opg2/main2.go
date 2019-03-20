package main

import (
	"flag"
	"ica03/opg2/fileinfo"
)

var file1, file2, file3 = "../files/text1.txt", "../files/text2.txt", "../files/pg100.txt"

func main() {
	f := flag.String("f", file1, "file to be used")
	fileinfo.PrintFileInfo(*f)
}
