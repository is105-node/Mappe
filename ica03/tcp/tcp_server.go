package main

import "net"
import "fmt"
import "bufio"
//import "encoding/json"
import "strings"
type NandEm struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

func main() {
	fmt.Println("Launching server...")


	l, _ := net.Listen("tcp", ":5000")
	
	conn, err := l.Accept()
	for {
			if err != nil {
			panic(err)
		}
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//inp := strings.Fields(message)
		//struc:= NandEm{inp[0], inp[1]}
		//nande, _ := json.Marshal(struc)
		nande:=""
		fmt.Println(message)
		if strings.Contains(message, "json") {
			conn.Write([]byte(string(nande) + "\n"))
		} else {
			fmt.Println("Received:", message)
			conn.Write([]byte("ok" + "\n"))		
		}
		
		}
}

