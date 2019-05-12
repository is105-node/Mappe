package main

import ("fmt"
		ft "fileutils"
		)

func main() {
	
	f1 := ft.FileToByteslice("files/text1.txt")
	f2 := ft.FileToByteslice("files/text2.txt")
	fmt.Print("Text 1: ", len(f1))
	fmt.Print("\n")
	fmt.Print("Text 2: ", len(f2))
}
