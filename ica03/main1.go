//C:\Users\Administrator\Documents\105\ica03\is105-ica03\files
package main

import ("os"
		"fmt"
		//"regexp"
		"io/ioutil"
		//"strings"
		)
func main () {

	args := os.Args
	filename:= args[1]
	LineShiftFinder(filename)
} 

func LineShiftFinder(filename string) string {
	
	f,err:= ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	var FeilOppstar [1]byte 
	FeilOppstar[0] = 92
	for _, e := range f {
		if (e != FeilOppstar[0]) {	
			fmt.Print(string(e))
		}
		

	}
	return ""
	
}