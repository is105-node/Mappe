package main

import(

      "fmt"
      "bufio"
      "os"
      "strings"
      "os/exec"
      )
//speech.SetWitKey("7CX3AO2H4EXLQ4WAWEQL54QHP54EUHPT")
//fmt.Println(speech.SendWitMessage("Test line"))
func main(){
  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Print("Press enter when you are ready to record. Type exit if you are done. \n>")
  input, _ := consoleReader.ReadString('\n')
  input = strings.ToLower(input)

  rec, _ := exec.Command("./gspeech-rec.sh").Output()
  if strings.Index(string(rec),"confidence:") == -1 {
    fmt.Println("No audio recognized.")
  } else {
  tsk := strings.Index(string(rec), "confidence:")
  cr := string(rec[51:tsk-1])
  fmt.Println(cr)
  }


}
