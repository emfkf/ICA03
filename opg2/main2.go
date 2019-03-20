package main

import (
	"ica03/opg2/fileinfo"
)

var file1, file2, file3 = "../files/text1.txt", "../files/text2.txt", "../files/pg100.txt"

func main() {
	fileinfo.PrintFileInfo(file1)
}
