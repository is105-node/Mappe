package main

import (
	"fmt"
	"fileutils"
	)

func main(){    
	as:=fileutils.FileToByteslice("files/lang01.wl")
	bs:=fileutils.FileToByteslice("files/lang02.wl")
	cs:=fileutils.FileToByteslice("files/lang03.wl")
	fmt.Printf("% c \n\n", as[:30])
	fmt.Printf("% c \n\n", bs[:30])
	fmt.Printf("% c \n\n", cs[:30])
}