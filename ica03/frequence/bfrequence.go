package bfrequence
	
import ("fmt"
		"bufio"
		"os"
		"io/ioutil"
		"strings"
	//	"bytes"
		"strconv"
		//"unicode/utf8"
	)


func main() {
    args := os.Args
	filename := args[2]
	lines := fmt.Sprintf("Lines in the file: %v\r\n", CountLines(filename))
	data := CountRunes(filename)
	fn:= fileNamer()
	data = lines + data
	WriteFile(fn, data)
}

func WriteFile(fn string, info string){
	fn = "./tests/" + fn
    file, err := os.Create(fn)
    if err != nil {
    	fmt.Println(err)
    } else {
    buffWrite := bufio.NewWriter(file)
    buffWrite.WriteString(info)
    buffWrite.Flush()
    
	}
	file.Close()
}


func fileNamer() string{
	var fileName string
	dir, _ := ioutil.ReadDir("./tests/")
    s:= []string{"bfrequence_res", strconv.Itoa(len(dir) + 1), ".txt"}
    fileName = strings.Join(s, "")
	return fileName

}



func CountLines(filename string) string{

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	count := 0
	var txt []string
	
	for scanner.Scan() {
		txt = append(txt, scanner.Text())
		count++
	}
	file.Close()

	return fmt.Sprintln("Lines: ", count)

	

}

func CountRunes(filename string) string{
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var txt []string
	
	for scanner.Scan() {
		txt = append(txt, scanner.Text())
	}

	var info string
	run := make(map[string] int) 

	for _, sym := range txt {
		char := sym
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
		 info = info + fmt.Sprintf("Rune: %s - %v Instances\r\n", top5lett[x], top5Val[x])
	}
	
	return info
	
}