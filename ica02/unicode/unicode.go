package main

import (
  "fmt"
  "os"
  "strings"
)
//UTF8 is the way to go
func main() {
   args := os.Args
   ln := args[1]
   fmt.Println(Translate("nord og sør", strings.ToUpper(ln)))
   UnicodeCodePointDemo()
}



// Kode for Oppgave 4a
func Translate(expression string, language string) string {
	var ret string
  if language=="IS" {
    expression = "“norður og suður”"
  }
  if language == "JP" {
    expression = "“北と南”"
  }
  for x:=0;x<len(expression);x++ {
    ret+=fmt.Sprintf("\\x%X", expression[x])
  }
  fmt.Printf("%s\n", expression)
  return ret
}

// Kode for Oppgave 4b
func UnicodeCodePointDemo() {
	// Hva er dette?
	// Er det likt på MS Windows og macOS?
  fmt.Println("\xf0\x9F\x98\x80")
  fmt.Println("\xf0\x9F\x98\x97")
  // Demonstrerer at deler av et tegn representert med flere bytes
  // kan ikke tolkes innenfor koden (unicode)
  fmt.Println("\xf0\x9F\x98")
  fmt.Println("\xf0\x9F")
  fmt.Println("\xf0")
}

