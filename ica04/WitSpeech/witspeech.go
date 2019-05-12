package main

import("github.com/nicolaifsf/go-speak"
      "fmt"
      "bufio"
      "os"
      "strings"
      )
//speech.SetWitKey("7CX3AO2H4EXLQ4WAWEQL54QHP54EUHPT")
//fmt.Println(speech.SendWitMessage("Test line"))

func main(){
  speech.SetWitKey("UOH2IONZ33BVSQPHTDETA63SI3SYKNM2")

  consoleReader := bufio.NewReader(os.Stdin)
  fmt.Print("Hit enter.\n>")
  _, err := consoleReader.ReadString('\n')
  if (err != nil) {
    fmt.Print("Error")
  }

  interpreted := speech.SendWitVoice("record_1.wav")
  inx := strings.Index(interpreted[15:], "\",")

  fmt.Println(inx)
  fmt.Println("You said: ", interpreted[15:15+inx])
}
