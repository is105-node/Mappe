package main

import ("pipe"
		"os"
		"fmt")



func main() {
	args := os.Args
	if len(args) <= 2 {
		os.Stderr.WriteString("Syntax to run: <filename> <utf8/base64/gzip> without <>")
	} else {
	args =args[1:]
	switch args[1] {
	case "utf8":
		f:= pipe.Reputf8(args[0])
		fmt.Printf("%s\nBytes Written: %v", f, len(f))
	case "base64":
		f:= pipe.Repb64(args[0])
		fmt.Printf("%s\nBytes Written: %v", f, len(f))
	case "gzip":
		pipe.Repgzip(args[0])
	}

	}


}