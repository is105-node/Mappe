package frequence
	
import ("fmt"
		"io/ioutil"
		"os"
		"bytes"
		"strconv"
		"strings"
		//"unicode/utf8"
	)


func main() {
    args := os.Args
	filename := args[2]
	dat, err := ioutil.ReadFile(filename)
    if err != nil {
    	fmt.Println(err)
    }
	st := fmt.Sprintf("Lines in the file: %v\r\n", CountLines(dat))
	st = st+ CountRunes(string(dat))
	fn:= fileNamer()
	err = writeFile(fn, st)
	if err != nil {
		fmt.Println("An error occured while writing to file.")
	} else {
		fmt.Println("Written to file ", fn)
	}
}

func writeFile(fn string, info string) error{
    fn = "./tests/"+fn
    newFile, _ := os.Create(fn)
	err := ioutil.WriteFile(fn, []byte(info), 7777)
	newFile.Close()
	return err
}


func fileNamer() string{
	var fileName string
	dir, _ := ioutil.ReadDir("./tests/")
    s:= []string{"frequence_res", strconv.Itoa(len(dir) + 1), ".txt"}
    fileName = strings.Join(s, "")
	return fileName
}

func CountLines(file []byte) int{
	newLine := []byte{'\n'}	
	c:= bytes.Count(file, newLine)
	return c
}

func CountRunes(file string) string{
	var toWr string
	run := make(map[string] int) 

	for _, sym := range file {
		char := strconv.QuoteRune(sym)
		count := run[char]

		if count == run[char] {
			run[char] = count + 1
		} else {
			run[char] = 1
		}
	}
	top5lett := [5]string{"", "", "", "", ""}
	top5Val := [5]int{0,0,0,0,0}

	for k, v := range run {

		switch val := v;  {
		case val > top5Val[0]:
			top5Val[0] = val
			top5lett[0] = k
		case val > top5Val[1]:
			top5Val[1] = val
			top5lett[1] = k
		case val > top5Val[2]:	
			top5Val[2] = val
			top5lett[2] = k
		case val > top5Val[3]:
			top5Val[3] = val
			top5lett[3] = k
		case val > top5Val[4]:
			top5Val[4] = val
			top5lett[4] = k
		}
	}
	
	for x:=0; x<len(top5Val);x++{
		toWr = toWr + fmt.Sprintf("Rune: %s - %v Instances\r\n", top5lett[x], top5Val[x])
	}
	
	return toWr
	
}