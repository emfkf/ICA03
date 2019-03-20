package lineshift

import (
	"fmt"
	"ica03/fileutils"
)

// GetNewLineTypeUsed ...
func GetNewLineTypeUsed(filename string) string {
	fileSlice := fileutils.FileToByteSlice(filename)
	newLineType := "No newlines"

	for i := 0; i < len(fileSlice)-1; i++ {
		if fmt.Sprintf("%X", fileSlice[i]) == "A" {
			newLineType = "LF"
		} else if fmt.Sprintf("%X", fileSlice[i]) == "D" && fmt.Sprintf("%X", fileSlice[i+1]) == "A" {
			newLineType = "CR LF"
		}
	}

	return newLineType
}
