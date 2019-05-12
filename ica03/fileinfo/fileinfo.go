package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	fmt.Println(len(args))
	if len(args) <= 2 {
		fmt.Println("Use syntax -f filename")
		return
	} 
	filename := args[2]
	fmt.Println(filename)
	fi, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}
	abPath, _ := filepath.Abs(filename)
	fmt.Println("Information about a file "+abPath+":")
	fmt.Println("Size: ", fi.Size() , " bytes, ", float32(fi.Size()) / float32(1024), " kibibytes, ", float32(fi.Size()) / float32(1049000), " mebibytes, ", float32(fi.Size()) / float32(1074000000), " gigibytes")
	switch file:=fi.Mode(); {
	case file.IsDir():
		fmt.Println("Is a directory")
	case !file.IsDir():
		fmt.Println("Is not a directory")
	}
	switch file:=fi.Mode(); {
	case file.IsRegular():
		fmt.Println("Is a regular file")
	case !file.IsRegular():
		fmt.Println("Is not regular file")
	}
	fmt.Println("Has Unix permission bits: ", fi.Mode().Perm())
	switch file:=fi.Mode(); {
	case file&os.ModeAppend == 1:
		fmt.Println("Is append only")
	case file&os.ModeAppend != 1:
		fmt.Println("Is not append only")
	}
	switch file:=fi.Mode(); {
	case file&os.ModeDevice == 1:
		fmt.Println("Is a device file")
	case file&os.ModeDevice != 1:
		fmt.Println("Is not a device file")
	}
	switch file:=fi.Mode(); {
	case file&os.ModeCharDevice != os.ModeCharDevice:
		fmt.Println("Is a Unix character device")
	case file&os.ModeCharDevice == os.ModeCharDevice:
		fmt.Println("Is not a Unix character device")
	}
	switch file:=fi.Mode(); {
	case file&os.ModeDevice == 1:
		fmt.Println("Is a block device")
	case file&os.ModeDevice != 1:
		fmt.Println("Is not a block device")
	}
	switch file:=fi.Mode(); {
	case file&os.ModeSymlink == 1:
		fmt.Println("Is a symbolic link")
	case file&os.ModeSymlink != 1:
		fmt.Println("Is not a symbolic link")
	}
}