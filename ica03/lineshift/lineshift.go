package main

import ("fmt"
		"os"
		ft "fileutils"
	)

func main () {
	AnalyzeNewline()
}

func AnalyzeNewline()  {
	args := os.Args
	filename := args[1]
	fileBS := ft.FileToByteslice(filename)
	newline:= "No new line."

	for i:=0; i < len(fileBS) - 1 ; i++ {
		if (fileBS[i] == 10) {
			newline = "Line feed only, (Mac/Linux) - n only."
			j:=i-1
		if (fileBS[j] == 13) {
			newline = "Carriage return + new line (Windows) r and n."
			}
		}
	}
	fmt.Println(newline)
}