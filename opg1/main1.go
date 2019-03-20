package main

import (
	"fmt"
	"ica03/fileutils"
	"ica03/lineshift"
)

// Path to files text1.txt and text2.txt
var file1, file2, file3 = "../files/text1.txt", "../files/text2.txt", "../files/pg100.txt"

func main() {
	testOpg1()
}

func testOpg1() {
	fileSlice1 := fileutils.FileToByteSlice(file1)
	fileSlice2 := fileutils.FileToByteSlice(file2)

	// Make sure that the max index is set to the file with the lowest length
	maxIndex := len(fileSlice1)
	if len(fileSlice1) > len(fileSlice2) {
		maxIndex = len(fileSlice2)
	}

	// Max index is one less because we're printing one value above the index
	for i := 0; i < maxIndex-1; i++ {
		if fileSlice1[i] != fileSlice2[i] {
			// Prints out the ASCII Hexadecimal of the 2 next bytes where the files
			// stop being identical
			fmt.Printf("File 1[%v] ASCII HEX: %X \n", i, fileSlice1[i])
			fmt.Printf("File 2[%v] ASCII HEX: %X \n", i, fileSlice2[i])
			fmt.Printf("File 1[%v] ASCII HEX: %X \n", i+1, fileSlice1[i+1])
			fmt.Printf("File 2[%v] ASCII HEX: %X \n", i+1, fileSlice2[i+1])
			// Break because we have all the information needed regarding these 2 files
			break
		}
	}

	str1 := lineshift.GetNewLineTypeUsed(file1)
	str2 := lineshift.GetNewLineTypeUsed(file2)
	str3 := lineshift.GetNewLineTypeUsed(file3)

	fmt.Println("File 1 uses: " + str1)
	fmt.Println("File 2 uses: " + str2)
	fmt.Println("File 3 uses: " + str3)
}
