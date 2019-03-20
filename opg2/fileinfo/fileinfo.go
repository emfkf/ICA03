package fileinfo

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

// PrintFileInfo prints information about a given file
func PrintFileInfo(filename string) {
	fileInfo, err = os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Information about a file", fileInfo.Name(), ":")
	fmt.Println("Size", fileInfo.Size(), "in bytes")
	fmt.Println("Is a directory:", fileInfo.IsDir())
	fmt.Println("Is regular:", fileInfo.Mode().IsRegular())
	fmt.Println("Has Unix permission bits:", fileInfo.Mode().Perm())

	isAppendOnly, isDeviceFile, isUnixCharDev, isUnixBlockDev, isSymbolicLink := false, false, false, false, false
	switch mode := fileInfo.Mode(); {
	case mode&os.ModeExclusive != 0:
		isAppendOnly = true
	case mode&os.ModeDevice != 0:
		isDeviceFile = true
		if mode%os.ModeCharDevice != 0 {
			isUnixCharDev = true
		} else {
			isUnixBlockDev = true
		}
	case mode&os.ModeSymlink != 0:
		isSymbolicLink = true
	}
	fmt.Println("Is append only: ", isAppendOnly)
	fmt.Println("Is a device file:", isDeviceFile)
	fmt.Println("Is a Unix character device:", isUnixCharDev)
	fmt.Println("Is a Unix block device: ", isUnixBlockDev)
	fmt.Println("Is a symbolic link:", isSymbolicLink)
}
